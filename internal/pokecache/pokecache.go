package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]CacheEntry
	mu       sync.Mutex
	interval time.Duration
	ticker   *time.Ticker
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(timeDur time.Duration) *Cache {
	cache := Cache{
		cacheMap: make(map[string]CacheEntry),
		interval: timeDur,
		ticker:   time.NewTicker(timeDur),
	}

	go cache.reapLoop()

	return &cache
}

func (c *Cache) Add(key string, byteVal []byte) {
	c.mu.Lock()
	newCacheEntry := CacheEntry{
		createdAt: time.Now(),
		val:       byteVal,
	}
	c.cacheMap[key] = newCacheEntry
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, ok := c.cacheMap[key]
	c.mu.Unlock()
	if !ok {
		return []byte{}, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() {
	defer c.ticker.Stop()

	for range c.ticker.C {
		// fmt.Printf("Interval passed at %s! Running reap loop.\n", time.Now().Format("15:04:05.000000"))
		c.mu.Lock()
		for key, val := range c.cacheMap {
			t1 := time.Now()
			createdDuration := t1.Sub(val.createdAt)
			if createdDuration > c.interval {
				// fmt.Printf("Interval passed! %v value deleted.\n", key)
				delete(c.cacheMap, key)
			}
		}
		c.mu.Unlock()
	}
}
