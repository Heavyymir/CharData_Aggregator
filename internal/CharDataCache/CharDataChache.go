package CharDataCache

import (
	"sync"
	"time"
)

// Define cacheEntry Data. Formats API responses for the Cache.
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Define Cache Data. Stores cached API responses.
type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

// NewCache initialises an empty cache and starts the cleanup subroutine
func NewCache(interval time.Duration) *Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}

	go c.reapLoop(interval)
	return &c
}

// Adds a cacheEntry to the Cache
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Gets the cacheEntry at a specified key
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

// Sets a timed subroutine to check and clean the cache at an interval by deleting old entries.
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.cache {
			if time.Since(entry.createdAt) > interval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}
