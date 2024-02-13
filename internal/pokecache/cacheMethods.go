package pokecache

import (
	"time"
)

func (c *Cache) Add(k string, v []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[k] = cacheEntry{ 
		createdAt: time.Now(),	
		val: v,
	}
}

func (c *Cache) Get(k string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.cache[k]
	if !ok {
		return nil, ok
	}
	return v.val, ok
}

func (c *Cache) reapLoop(i time.Duration) {
	ticker := time.NewTicker(i * time.Second)	

	go func() {
		for {
			select {
			case t := <-ticker.C:
				for k, v := range c.cache {
					if t.Add(i * time.Second).After(v.createdAt) {
						c.mu.Lock()
						delete(c.cache, k)
						c.mu.Unlock()
					}	
				}
			}
		}
	}()
}
