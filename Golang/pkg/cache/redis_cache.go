package cache

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"

	"github.com/elberthcabrales/movies-api/pkg/config"
)

var ctx = context.Background()

// RedisCache is a wrapper around the Redis client providing caching functionality.
type RedisCache struct {
	client *redis.Client
}

// NewRedisCache creates a new RedisCache using the provided Redis client and configuration.
func NewRedisCache(client *redis.Client, cfg *config.RedisConfig) *RedisCache {
	if cfg != nil {
		client = redis.NewClient(&redis.Options{
			Addr:     cfg.Addr,
			Password: cfg.Password,
			DB:       cfg.DB,
		})

		log.Println("Pinging Redis...")
		_, err := client.Ping(ctx).Result()
		if err != nil {
			log.Fatalf("Failed to connect to Redis: %v", err)
		}
		log.Println("Successfully connected to Redis")
	}

	return &RedisCache{
		client: client,
	}
}

// SetValue sets a key-value pair in the Redis cache.
func (r *RedisCache) SetValue(key string, value interface{}) error {
	log.Printf("Setting value in Redis for key: %s", key)
	err := r.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Printf("Failed to set value for key %s: %v", key, err)
		return err
	}
	log.Printf("Value set successfully for key: %s", key)
	return nil
}

// GetValue retrieves the value associated with the key from the Redis cache.
func (r *RedisCache) GetValue(key string) (string, error) {
	log.Printf("Getting value from Redis for key: %s", key)
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		log.Printf("Failed to get value for key %s: %v", key, err)
		return "", err
	}
	log.Printf("Successfully retrieved value for key: %s", key)

	return val, nil
}
