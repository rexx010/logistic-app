package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"logisticApp/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const idempotencyTTL = 24 * time.Hour

type storedResponse struct {
	StatusCode int             `json:"status_code"`
	Body       json.RawMessage `json:"body"`
}

func Idempotency() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method == http.MethodGet || method == http.MethodHead || method == http.MethodOptions {
			c.Next()
			return
		}

		idempotencyKey := strings.TrimSpace(c.GetHeader("Idempotency-Key"))
		if idempotencyKey == "" {
			c.Next()
			return
		}

		if len(idempotencyKey) > 128 {
			utils.BadRequest(c, "Idempotency-Key must be 128 characters or fewer", nil)
			c.Abort()
			return
		}

		ctx := context.Background()
		redisKey := utils.CacheKeyIdempotency + idempotencyKey
		var cached storedResponse
		found, err := utils.Get(ctx, redisKey, &cached)
		if err == nil && found {
			c.Header("X-Idempotent-Replayed", "true")
			c.Data(cached.StatusCode, "application/json", cached.Body)
			c.Abort()
			return
		}

		recorder := &responseRecorder{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
			statusCode:     http.StatusOK,
		}
		c.Writer = recorder
		c.Next()
		responseBody := recorder.body.Bytes()
		if len(responseBody) > 0 {
			toCache := storedResponse{
				StatusCode: recorder.statusCode,
				Body:       json.RawMessage(responseBody),
			}
			_ = utils.Set(ctx, redisKey, toCache, idempotencyTTL)
		}
	}
}

type responseRecorder struct {
	gin.ResponseWriter
	body       *bytes.Buffer
	statusCode int
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func RequireIdempotencyKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		Key := strings.TrimSpace(c.GetHeader("Idempotency-Key"))
		if Key == "" {
			utils.BadRequest(c, "Idempotency-Key is required for this endpoint", nil)
			c.Abort()
			return
		}
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		c.Next()
	}
}
