package pokecache

import (
	"time"
)

var MemCache *Cache

func InitGlobalCache(interval time.Duration) {
	if MemCache == nil {
		MemCache = NewCache(interval)
	}
}
