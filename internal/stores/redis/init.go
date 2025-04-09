package redisstore

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// NewRedisCache ...
func NewRedisCache(client *redis.Client, expiration time.Duration) *RedisCache {
	if expiration <= 0 {
		expiration = time.Hour
	}

	redisCache := &RedisCache{
		client:     client,
		ctx:        context.Background(),
		expiration: expiration,
	}

	if err := redisCache.Ping(); err != nil {
		panic("Redis Connection error" + err.Error())
	}
	return redisCache
}
