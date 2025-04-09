package memorystore

import (
	"time"
)

func (c *LRUCache) evict() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.order.Len() == 0 {
		return
	}

	element := c.order.Back()
	if element == nil {
		return
	}

	node := element.Value.(*CacheNode)
	delete(c.memory, node.key)
	c.order.Remove(element)
}

// StartExpireCycle ...
func (c *LRUCache) StartExpireCycle() {
	c.evictionPolicy.signal = make(chan struct{})

	go c.expireCycle()
}

// StopExpireCycle ...
func (c *LRUCache) StopExpireCycle() {
	if c.evictionPolicy.signal != nil {
		close(c.evictionPolicy.signal)
		c.evictionPolicy.signal = nil
	}
}

// expireCycle ...
func (c *LRUCache) expireCycle() {
	// logger := helpers.Log
	ticker := time.NewTicker(c.evictionPolicy.interval)
	defer ticker.Stop()

	stopping := false

	for {
		if stopping {
			return
		}
		select {
		case <-ticker.C:
			c.runBatch()
		case <-c.evictionPolicy.signal:
			stopping = true
			ticker.Stop()
		}
	}
}

// runBatch ...
func (c *LRUCache) runBatch() {
	// logger := helpers.Log

	c.mu.Lock()
	defer c.mu.Unlock()

	element := c.evictionPolicy.lastScan

	if element == nil {
		element = c.order.Front()
	}

	count := 0
	now := time.Now()

	for element != nil && count < c.evictionPolicy.batchSize {
		next := element.Next()
		node := element.Value.(*CacheNode)
		data := node.data

		if !data.forever && now.After(data.expiredAt) {
			delete(c.memory, node.key)
			c.order.Remove(element)
		}

		element = next
		count++
	}
	c.evictionPolicy.lastScan = element
	if c.evictionPolicy.lastScan == nil {
		c.evictionPolicy.lastScan = c.order.Front()
	}
}
