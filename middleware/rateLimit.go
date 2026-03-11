package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"logisticApp/config"
	"logisticApp/utils"

	"github.com/gin-gonic/gin"
)

// RateLimitConfig defines the rules for a rate limiter.
type RateLimitConfig struct {
	// Max number of requests allowed in the Window.
	Limit int
	// The time window (e.g. 1 minute).
	Window time.Duration
	// KeyFunc determines HOW to identify the caller.
	// Default: by IP. Can be overridden to limit by user ID.
	KeyFunc func(c *gin.Context) string
}

// RateLimit returns a Redis-based sliding window rate limiter middleware.
//
// Why Redis and not in-memory?
// With a load balancer sending requests across 3 instances:
//   - In-memory: each instance counts separately → user gets 3x the limit
//   - Redis: all instances share one counter → limit is truly enforced
//
// We use the "fixed window counter" algorithm — simple and effective.
// For more precision you'd use "sliding window log" (more Redis memory usage).
func RateLimit(cfg RateLimitConfig) gin.HandlerFunc {
	// Default: identify callers by their IP address.
	if cfg.KeyFunc == nil {
		cfg.KeyFunc = func(c *gin.Context) string {
			// c.ClientIP() respects X-Forwarded-For header set by Nginx,
			// so we get the real client IP, not the load balancer's IP.
			return c.ClientIP()
		}
	}

	return func(c *gin.Context) {
		ctx := context.Background()
		caller := cfg.KeyFunc(c)

		// Build a time-bucketed key.
		// We divide time into fixed windows (buckets) based on the current time.
		// e.g. for a 1-minute window: bucket = Unix timestamp / 60
		// All requests within the same minute share the same Redis key.
		bucket := time.Now().Unix() / int64(cfg.Window.Seconds())
		redisKey := fmt.Sprintf("%s%s:%d", utils.CacheKeyRateLimit, caller, bucket)

		// INCR atomically increments the counter and returns the new value.
		// Atomic = thread-safe, race-condition-free, even across multiple instances.
		count, err := config.RedisClient.Incr(ctx, redisKey).Result()
		if err != nil {
			// If Redis is down, fail open (allow the request) rather than
			// blocking all users. Adjust this policy based on your risk tolerance.
			c.Next()
			return
		}

		// Set the TTL on the first request of each window.
		// We use 2x the window so the key doesn't expire mid-window.
		if count == 1 {
			config.RedisClient.Expire(ctx, redisKey, cfg.Window*2)
		}

		// Set rate limit headers so clients know their current status.
		// This is the standard pattern used by GitHub, Stripe, etc.
		remaining := cfg.Limit - int(count)
		if remaining < 0 {
			remaining = 0
		}
		c.Header("X-RateLimit-Limit", strconv.Itoa(cfg.Limit))
		c.Header("X-RateLimit-Remaining", strconv.Itoa(remaining))
		c.Header("X-RateLimit-Reset", strconv.FormatInt((bucket+1)*int64(cfg.Window.Seconds()), 10))

		if int(count) > cfg.Limit {
			c.JSON(http.StatusTooManyRequests, utils.APIResponse{
				Success: false,
				Message: "Too many requests. Please slow down.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// ── Pre-configured rate limiters for common use cases ────────────────

// GlobalRateLimit applies a broad limit to all routes.
// Protects against general abuse / DDoS.
func GlobalRateLimit() gin.HandlerFunc {
	return RateLimit(RateLimitConfig{
		Limit:  300,         // 300 requests...
		Window: time.Minute, // ...per minute, per IP
	})
}

// AuthRateLimit is stricter — applied only to login/register routes.
// Prevents brute-force password attacks.
func AuthRateLimit() gin.HandlerFunc {
	return RateLimit(RateLimitConfig{
		Limit:  10,          // 10 attempts...
		Window: time.Minute, // ...per minute, per IP
	})
}

// UserRateLimit limits by authenticated user ID instead of IP.
// Use this on expensive endpoints (e.g. payment initiation).
// Prevents a single user from hammering an endpoint even behind a proxy.
func UserRateLimit(limit int, window time.Duration) gin.HandlerFunc {
	return RateLimit(RateLimitConfig{
		Limit:  limit,
		Window: window,
		KeyFunc: func(c *gin.Context) string {
			// "userID" is set by the JWT Auth middleware (coming next step).
			if userID, exists := c.Get("userID"); exists {
				return fmt.Sprintf("user:%v", userID)
			}
			return c.ClientIP() // fallback to IP if not authenticated
		},
	})
}
