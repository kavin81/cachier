package memorystore

import (
	"time"
)

// Get ...
func (c *LRUCache) Get(key string) (string, error) {
	c.StopExpireCycle()
	defer c.StartExpireCycle()

	node, exists := c.safeLoad(key)
	if exists {
		return node.data.value, nil
	}
	// TODO: custom error  implementation
	return "", nil
}

// Set ...
func (c *LRUCache) Set(key string, value string) error {
	// logger := helpers.Log

	c.StopExpireCycle()
	defer c.StartExpireCycle()

	// update
	if node, exists := c.safeLoad(key); exists {
		node.data.value = value
		return nil
	}

	// check if the cache is full
	if c.order.Len() >= c.capacity {
		c.evict()
	}

	// new
	elem := c.order.PushFront(&CacheNode{
		key: key,
		data: &CacheData{
			value:     value,
			forever:   false,
			expiredAt: time.Now().Add(c.evictionPolicy.expiration),
		},
	})
	c.memory[key] = elem

	return nil
}

// Delete ...
func (c *LRUCache) Delete(key string) (int, error) {
	c.StopExpireCycle()
	c.mu.Lock()
	defer c.StartExpireCycle()

	element, ok := c.memory[key]
	if !ok {
		c.mu.Unlock()
		return 0, nil
	}

	node := element.Value.(*CacheNode)
	data := node.data

	c.order.Remove(element)
	delete(c.memory, key)
	if !data.forever && time.Now().After(data.expiredAt) {
		c.mu.Unlock()
		// TODO: custom error  implementation
		return 0, nil
	}

	c.mu.Unlock()
	return 1, nil
}
