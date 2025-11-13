package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Locations map[string]cacheEntry
	mu        sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// expose a NewCache() function that creates a new cache with a configurable interval (time.Duration)
func NewCache(Interval time.Duration) *Cache {
	c := &Cache{
		Locations: make(map[string]cacheEntry),
	}
	go c.reapLoop(Interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Locations[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, exists := c.Locations[key]
	if exists {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.Locations {
			if time.Since(entry.createdAt) > interval {
				delete(c.Locations, key)
			}
		}
		c.mu.Unlock()
	}
}
