package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	CacheMap map[string]CacheEntry
	Mu       sync.Mutex
	Interval time.Duration
	ticker   *time.Ticker
}

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

func NewCache(timeDur time.Duration) *Cache {
	cache := Cache{
		CacheMap: make(map[string]CacheEntry),
		Interval: timeDur,
		ticker:   time.NewTicker(timeDur),
	}

	go cache.reapLoop()

	return &cache
}

func (c *Cache) Add(key string, byteVal []byte) {
	newCacheEntry := CacheEntry{
		CreatedAt: time.Now(),
		Val:       byteVal,
	}
	c.CacheMap[key] = newCacheEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.CacheMap[key]

	if !ok {
		return []byte{}, false
	}

	return entry.Val, true
}

func (c *Cache) reapLoop() {
	defer c.ticker.Stop()

	for range c.ticker.C {
		for key, val := range c.CacheMap {
			t1 := time.Now()
			createdDuration := t1.Sub(val.CreatedAt)
			if createdDuration > c.Interval {
				delete(c.CacheMap, key)
			}
		}
	}
}
