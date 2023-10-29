package cache

import (
	"errors"
	"sync"
	"time"
)

type item struct {
	value interface{}
	exp   int64
}

type Cache struct {
	cache map[interface{}]*item
	sync.RWMutex
}

func NewCache() *Cache {
	cache := &Cache{
		cache: make(map[interface{}]*item),
	}
	go cache.cleaner()
	return cache
}

func (c *Cache) cleaner() {
	for {
		c.Lock()
		for k, v := range c.cache {
			if time.Now().Unix() > v.exp {
				delete(c.cache, k)
			}
		}
		c.Unlock()
		<-time.After(1 * time.Second)
	}
}

func (c *Cache) Set(key interface{}, value interface{}) error {
	c.Lock()
	defer c.Unlock()

	c.cache[key] = &item{
		value: value,
		exp:   time.Now().Unix() + 500,
	}

	return nil
}

func (c *Cache) Get(key interface{}) (interface{}, error) {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.cache[key]; ok {
		return c.cache[key], nil
	}

	return nil, errors.New("item not found")
}
