package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// RedisConfig holds the configuration settings for connecting to Redis.
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

// LoadConfig loads environment variables and returns a RedisConfig struct
func LoadConfig() *RedisConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	return &RedisConfig{
		Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
		Password: getEnv("REDIS_PASSWORD", ""),
		DB:       getEnvAsInt("REDIS_DB", 0),
	}
}

// Helper functions to read environment variables
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(name string, defaultValue int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
