package memorystore

import "container/list"

// DeleteMany ...
func (c *LRUCache) DeleteMany(keys []string) (int, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	count := 0

	for _, key := range keys {
		tmpCount, _ := c.Delete(key)
		count += tmpCount
	}
	return count, nil
}

// Exists ...
func (c *LRUCache) Exists(key string) bool {
	c.StopExpireCycle()
	defer c.StartExpireCycle()

	_, exists := c.safeLoad(key)
	return exists
}

// Pop ...
func (c *LRUCache) Pop(key string) (string, error) {
	if node, exists := c.safeLoad(key); exists {
		tmpData := node.data.value
		c.Delete(key)
		return tmpData, nil
	}
	// TODO: custom error  implementation
	return "", nil
}

// Flush ...
func (c *LRUCache) Flush() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.memory = make(map[string]*list.Element)
	c.order.Init()
	return nil
}
