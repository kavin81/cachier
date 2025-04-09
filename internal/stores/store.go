package stores

import (
	memorystore "github.com/kavin81/cachier/internal/stores/memory"
	redisstore "github.com/kavin81/cachier/internal/stores/redis"
)

// L1CacheSchema ...
type L1CacheSchema interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Delete(key string) (int, error)
	DeleteMany(keys []string) (int, error)
	Exists(key string) bool
	Pop(key string) (string, error)
	Flush() error
}

// L2CacheSchema ...
type L2CacheSchema interface {
	Ping() error
	L1CacheSchema
	Disconnect() error
}

var (
	// NewLRUCache ...
	NewLRUCache = memorystore.NewLRUCache
	// NewRedisCache ...
	NewRedisCache = redisstore.NewRedisCache
)
