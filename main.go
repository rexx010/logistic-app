package main

import (
	"log"
	"logisticApp/config"
	"logisticApp/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	config.ConnectDB()

	config.ConnectRedis()

	config.MigrateDB()

	if config.AppConfig.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.SetTrustedProxies([]string{"nginx", "127.0.0.1"})

	router.Use(middleware.GlobalRateLimit()) // 1st: drop abusive IPs fast
	router.Use(middleware.Idempotency())     // 2nd: replay duplicate requests

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "🚀 logisticApp is running"})
	})

	port := config.AppConfig.AppPort
	log.Printf("🚀 Server starting on port %s (env: %s)", port, config.AppConfig.AppEnv)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
