package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cEntry
	mux   *sync.Mutex
}

type cEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	ca := Cache{
		cache: make(map[string]cEntry),
		mux:   &sync.Mutex{},
	}
	go ca.reapLoop(interval)

	return ca
}

func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache[key] = cEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	v, ok := c.cache[key]
	return v.val, ok
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}
