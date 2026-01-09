package main

import (
	"time"

	"github.com/cardvark/pokedex/internal/pokecache"
)

func main() {
	pokecache.InitGlobalCache(5 * time.Second)
	startRepl()
}
