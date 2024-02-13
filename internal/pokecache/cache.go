package pokecache

import (
	"sync"
	"time"
)

func NewCache(i time.Duration) Cache {
	c := Cache{
		cache:	make(map[string]cacheEntry),
		mu:	&sync.Mutex{},	
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.reapLoop(i)
	return c 
}

