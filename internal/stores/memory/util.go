package memorystore

import (
	"time"

	"github.com/kavin81/cachier/internal/helpers"
	"go.uber.org/zap"
)

func (c *LRUCache) safeLoad(key string) (*CacheNode, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	element, ok := c.memory[key]
	if !ok {
		return nil, false
	}

	node := element.Value.(*CacheNode)
	data := node.data

	if !data.forever && time.Now().After(data.expiredAt) {
		c.order.Remove(element)
		delete(c.memory, key)
		return nil, false
	}

	c.order.MoveToFront(element)
	return node, true
}

// debug function
func (c *LRUCache) printAll() {
	logger := helpers.Log

	for key, node := range c.memory {
		logger.Info(zap.Any(key, node.Value.(*CacheNode).data.value))
	}
}
