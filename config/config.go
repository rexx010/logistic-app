package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	//Application
	AppEnv  string
	AppPort string

	//Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	//JWT
	JWTSecret    string
	JWTExpiresIn string

	//cloudinary
	CloudinaryCloudName string
	CloudinaryAPIKey    string
	CloudinaryAPISecret string

	//Termii
	TermiiAPIKey   string
	TermiiSenderID string
	TermiiBaseURL  string

	//Email
	SMTPHost     string
	SMTPPort     string
	SMTPUser     string
	SMTPPassword string
	SMTPFrom     string

	// Paystack
	PaystackSecretKey string
	PaystackBaseURL   string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int
}

var AppConfig *Config

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found — reading from system environment")
	}
	redisDB, _ := strconv.Atoi(getEnv("REDIS_DB", "0"))

	AppConfig = &Config{
		AppEnv:  getEnv("APP_ENV", "development"),
		AppPort: getEnv("APP_PORT", "8080"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "logistic_app"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),

		JWTSecret:    getEnvRequired("JWT_SECRET"),
		JWTExpiresIn: getEnv("JWT_EXPIRES_IN", "24h"),

		CloudinaryCloudName: getEnvRequired("CLOUDINARY_CLOUD_NAME"),
		CloudinaryAPIKey:    getEnvRequired("CLOUDINARY_API_KEY"),
		CloudinaryAPISecret: getEnvRequired("CLOUDINARY_API_SECRET"),

		TermiiAPIKey:   getEnvRequired("TERMII_API_KEY"),
		TermiiSenderID: getEnv("TERMII_SENDER_ID", "LogisticApp"),
		TermiiBaseURL:  getEnv("TERMII_BASE_URL", "https://api.ng.termii.com/api"),

		SMTPHost:     getEnv("SMTP_HOST", "smtp.gmail.com"),
		SMTPPort:     getEnv("SMTP_PORT", "587"),
		SMTPUser:     getEnvRequired("SMTP_USER"),
		SMTPPassword: getEnvRequired("SMTP_PASSWORD"),
		SMTPFrom:     getEnv("SMTP_FROM", "no-reply@logisticapp.com"),

		PaystackSecretKey: getEnvRequired("PAYSTACK_SECRET_KEY"),
		PaystackBaseURL:   getEnv("PAYSTACK_BASE_URL", "https://api.paystack.co"),

		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnv("REDIS_PORT", "6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       redisDB,
	}

	log.Println("✅ Config loaded successfully")
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Africa/Lagos",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode,
	)
}

func (c *Config) RedisAddr() string {
	return fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvRequired(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		log.Fatalf("❌ Required environment variable %q is not set", key)
	}
	return value
}
