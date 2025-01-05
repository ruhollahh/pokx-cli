package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu    *sync.Mutex
	store map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		mu:    &sync.Mutex{},
		store: make(map[string]cacheEntry),
	}

	go c.reapLoop(interval)

	return c
}
