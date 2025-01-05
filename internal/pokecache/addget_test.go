package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cases := []struct {
		url  string
		data []byte
	}{
		{
			url:  "https://example.com",
			data: []byte("testdata"),
		},
		{
			url:  "https://google.com",
			data: []byte("moretestdataa"),
		},
	}

	interval := 5 * time.Second
	cache := NewCache(interval)

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			cache.Add(c.url, c.data)
			data, exists := cache.Get(c.url)
			if !exists {
				t.Fatalf("expected to find key")
			}

			if string(data) != string(c.data) {
				t.Fatalf("expected to find data")
			}
		})
	}
}
