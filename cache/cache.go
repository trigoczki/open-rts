package cache

import (
	ruleset "open-rts/ruleset/resource"
	"sync"
)

var (
	once     sync.Once
	instance *Cache
)

type Cache struct {
	mu        sync.RWMutex
	resources []*ruleset.Resource
}

func GetInstance() *Cache {
	once.Do(func() {
		instance = &Cache{}
	})
	return instance
}

func (c *Cache) SetResources(resources []*ruleset.Resource) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.resources = resources
}

func (c *Cache) GetResources() []*ruleset.Resource {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.resources
}
