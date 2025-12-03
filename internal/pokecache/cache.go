package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntry map[string]cacheEntry
	mutex      *sync.Mutex
	interval   time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cacheEntry[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cacheEntry, ok := c.cacheEntry[key]
	if !ok {
		return []byte{}, false
	}
	return cacheEntry.val, ok
}

func (c *Cache) reapLoop() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for key, val := range c.cacheEntry {
		if time.Since(val.createdAt) > c.interval {
			delete(c.cacheEntry, key)
		}
	}
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cacheEntry: map[string]cacheEntry{},
		mutex:      &sync.Mutex{},
		interval:   interval,
	}
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			c.reapLoop()
		}
	}()

	return c
}
