package utils

import (
	"context"
	"encoding/json"
	"logisticApp/config"
	"time"
)

const (
	CasheKeyUser        = "user:"
	CasheKeyDelivery    = "delivery:"
	CasheKeyRiderList   = "rider:available"
	CacheKeyIdempotency = "idempotency:"
	CacheKeyRateLimit   = "rate_limit:"
)

func Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return config.RedisClient.Set(ctx, key, data, ttl).Err()
}

func Get(ctx context.Context, key string, dest interface{}) (bool, error) {
	data, err := config.RedisClient.Get(ctx, key).Bytes()
	if err != nil {
		if err.Error() == "redis: nil" {
			return false, nil
		}
		return false, err
	}
	if err := json.Unmarshal(data, dest); err != nil {
		return false, err
	}
	return true, nil
}

func delete(ctx context.Context, keys ...string) error {
	return config.RedisClient.Del(ctx, keys...).Err()
}

func Exists(ctx context.Context, key string) (bool, error) {
	count, err := config.RedisClient.Exists(ctx, key).Result()
	return count > 0, err
}

func SetNX(ctx context.Context, key string, value interface{}, ttl time.Duration) (bool, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return false, err
	}
	return config.RedisClient.SetNX(ctx, key, data, ttl).Result()
}

// SetString stores a plain string value in Redis with a TTL.
// Simpler than Set() — no JSON marshalling needed for primitive values.
func SetString(key string, value string, ttl time.Duration) error {
	return config.RedisClient.Set(context.Background(), key, value, ttl).Err()
}

// GetString retrieves a plain string value from Redis.
// Returns ("", false, nil) on cache miss.
func GetString(key string) (string, bool, error) {
	val, err := config.RedisClient.Get(context.Background(), key).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return "", false, nil
		}
		return "", false, err
	}
	return val, true, nil
}

// DeleteKey removes a single key from Redis without requiring a context.
// Convenience wrapper for simple fire-and-forget deletions.
func DeleteKey(key string) error {
	return config.RedisClient.Del(context.Background(), key).Err()
}
