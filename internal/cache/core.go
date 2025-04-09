package cache

import "github.com/kavin81/cachier/internal/stores"

// Options ...
type Options struct {
	L1 stores.L1CacheSchema
	L2 stores.L2CacheSchema
}

// UnifiedCache ...
type UnifiedCache struct {
	namespace string
	L1        stores.L1CacheSchema
	L2        stores.L2CacheSchema
}
