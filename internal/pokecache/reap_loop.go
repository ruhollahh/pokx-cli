package pokecache

import (
	"time"
)

func (c *Cache) reapLoop(interval time.Duration) {
	ch := time.Tick(interval)
	for range ch {
		c.mu.Lock()
		for key, val := range c.store {
			if val.createdAt.Before(time.Now().Add(-interval).UTC()) {
				delete(c.store, key)
			}
		}
		c.mu.Unlock()
	}
}
