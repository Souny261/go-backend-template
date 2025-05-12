package output

import (
	"context"
	"time"
)

// CacheRepository defines the output port for cache operations
type CacheRepository interface {
	
	// Get retrieves a value from the cache by key
	Get(ctx context.Context, key string) (string, error)
	
	// Set stores a value in the cache with an optional expiration time
	Set(ctx context.Context, key string, value string, expiration time.Duration) error
	
	// Delete removes a value from the cache by key
	Delete(ctx context.Context, key string) error
	
	// Exists checks if a key exists in the cache
	Exists(ctx context.Context, key string) (bool, error)
	
	// Increment increments a numeric value in the cache
	Increment(ctx context.Context, key string) (int64, error)
	
	// Expire sets an expiration time on a key
	Expire(ctx context.Context, key string, expiration time.Duration) error
}
