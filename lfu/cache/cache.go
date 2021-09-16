package cache

import "container/list"

func New(capacity int) *Cache {
	c := new(Cache)

	c.freqs = list.New()
	c.size = 0
	c.capacity = capacity
	c.byKey = make(map[string]*CacheItem)
	return c
}

//Set insert the given value to cache,
//if key already present update the entry with new value.
//otherwise create a new entry
func (c *Cache) Set(key string, value interface{}) {
	if item, ok := c.byKey[key]; ok {
		//if key found, update its value
		item.value = value
	} else {
		//new entry, create a cache item and insert into cache
		item := new(CacheItem)
		item.value = value
		item.key = key
		c.byKey[key] = item
		c.size++
	}
}

//Get get a value associated with the given key,
// if key not found return nil
func (c *Cache) Get(key string) interface{} {
	if item, ok := c.byKey[key]; ok {
		return item.value
	}
	return nil
}
