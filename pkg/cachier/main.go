// Package cachier ...
package cachier

import (
	"github.com/kavin81/cachier/internal/cache"
	"github.com/kavin81/cachier/internal/stores"
)

// Options ...
type Options struct {
	L1 stores.L1CacheSchema
	L2 stores.L2CacheSchema
}

var (
	// NewLRUCache ...
	NewLRUCache = stores.NewLRUCache
	// NewRedisCache ...
	NewRedisCache = stores.NewRedisCache
	// New ...
	New = cache.New
)
