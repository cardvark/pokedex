package main

import (
	"time"

	"github.com/cardvark/pokedex/internal/pokecache"
)

var MemCache = pokecache.NewCache(5 * time.Second)

func main() {
	startRepl()
}
