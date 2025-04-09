// Package memorystore ...
package memorystore

import (
	"container/list"
	"sync"
	"time"
)

// CacheData ...
type CacheData struct {
	value     string
	expiredAt time.Time
	forever   bool
}

// CacheNode ...
type CacheNode struct {
	key  string
	data *CacheData
}

// CacheEvictionPolicy ...
type CacheEvictionPolicy struct {
	expiration time.Duration
	interval   time.Duration
	signal     chan struct{}
	lastScan   *list.Element
	batchSize  int
}

// LRUCache ...
type LRUCache struct {
	capacity int
	memory   map[string]*list.Element
	order    *list.List

	mu             sync.RWMutex
	evictionPolicy CacheEvictionPolicy
}
