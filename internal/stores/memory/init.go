package memorystore

import (
	"container/list"
	"time"
)

// NewLRUCache ...
func NewLRUCache(capacity int, expiration time.Duration) *LRUCache {
	lruCache := &LRUCache{
		capacity: max(capacity, 1),
		memory:   make(map[string]*list.Element),
		order:    list.New(),

		evictionPolicy: CacheEvictionPolicy{
			expiration: expiration,
			// interval:   min(expiration/10, time.Minute),
			interval:  5 * time.Second,
			signal:    make(chan struct{}),
			batchSize: int(time.Minute.Seconds()),
		},
	}

	lruCache.StartExpireCycle()

	return lruCache
}
