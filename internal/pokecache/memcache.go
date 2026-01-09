package pokecache

import (
	"time"
)

var MemCache = NewCache(5 * time.Second)
