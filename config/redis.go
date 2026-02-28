package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         AppConfig.RedisAddr(),
		Password:     AppConfig.RedisPassword,
		DB:           AppConfig.RedisDB,
		PoolSize:     10,
		MinIdleConns: 3,
	})
	ctx := context.Background()
	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		log.Fatalf("failed to connect redis: %v", err)
	}

	log.Println("Redis connection established successfully")
}
