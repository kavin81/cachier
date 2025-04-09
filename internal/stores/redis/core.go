// Package redisstore ...
package redisstore

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisCache ...
type RedisCache struct {
	client     *redis.Client
	ctx        context.Context
	expiration time.Duration
}
