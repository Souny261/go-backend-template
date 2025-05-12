package redis

import (
	"backend/internal/core/ports/output"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Config holds the configuration for Redis
type Config struct {
	Host     string
	Port     string
	Password string
	DB       int
}

// RedisRepository implements the CacheRepository interface
type RedisRepository struct {
	client *redis.Client
}

// NewRedisRepository creates a new Redis repository
func NewRedisRepository(config Config) (*RedisRepository, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("❌ failed to connect to Redis: %w", err)
	}

	return &RedisRepository{client: client}, nil
}

// Get retrieves a value from the cache by key
func (r *RedisRepository) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

// Set stores a value in the cache with an optional expiration time
func (r *RedisRepository) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

// Delete removes a value from the cache by key
func (r *RedisRepository) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

// Exists checks if a key exists in the cache
func (r *RedisRepository) Exists(ctx context.Context, key string) (bool, error) {
	result, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return result > 0, nil
}

// Increment increments a numeric value in the cache
func (r *RedisRepository) Increment(ctx context.Context, key string) (int64, error) {
	return r.client.Incr(ctx, key).Result()
}

// Expire sets an expiration time on a key
func (r *RedisRepository) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return r.client.Expire(ctx, key, expiration).Err()
}

// Close closes the Redis client
func (r *RedisRepository) Close() error {
	return r.client.Close()
}

// Ensure RedisRepository implements CacheRepository
var _ output.CacheRepository = (*RedisRepository)(nil)
