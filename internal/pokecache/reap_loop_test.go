package pokecache

import (
	"testing"
	"time"
)

func TestReapLoop(t *testing.T) {
	interval := 5 * time.Millisecond
	waitTime := interval + 5*time.Millisecond
	url := "https://example.com"
	data := []byte("testdata")

	cache := NewCache(interval)
	cache.Add(url, data)
	_, exists := cache.Get(url)
	if !exists {
		t.Fatalf("expected to find key")
	}

	time.Sleep(waitTime)

	_, exists = cache.Get(url)
	if exists {
		t.Fatalf("expected to not find key")
	}
}
