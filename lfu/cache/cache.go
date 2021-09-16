package cache

import "container/list"

func (c *Cache) increment(item *CacheItem) {
	currentFrequency := item.frequencyParent

	var nextFrequencyCount int
	var nextFrequency *list.Element

	if currentFrequency == nil {
		//next frequency value will be 1 and whatever is at the front of the list becomes the next frequency node
		nextFrequencyCount = 1
		nextFrequency = c.freqs.Front()
	} else {
		nextFrequencyCount = currentFrequency.Value.(*FrequencyItem).freq + 1
		nextFrequency = currentFrequency.Next()
	}

	if nextFrequency == nil || nextFrequency.Value.(*FrequencyItem).freq != nextFrequencyCount {
		nextFrequencyItem := new(FrequencyItem)
		nextFrequencyItem.freq = nextFrequencyCount
		nextFrequencyItem.entries = make(map[*CacheItem]byte)

		if currentFrequency == nil {
			//push the new frequency entry to the front of the list
			nextFrequency = c.freqs.PushFront(nextFrequencyItem)
		} else {
			nextFrequency = c.freqs.InsertAfter(nextFrequencyItem, currentFrequency)
		}
	}
	item.frequencyParent = nextFrequency
	//add entry for the item in the frequnecy list
	nextFrequency.Value.(*FrequencyItem).entries[item] = 1

	//remove the entry from the current frequency, as it is moved to new next entry
	if currentFrequency != nil {
		delete(currentFrequency.Value.(*FrequencyItem).entries, item)
	}
}

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
		c.increment(item)
	} else {
		//new entry, create a cache item and insert into cache
		item := new(CacheItem)
		item.value = value
		item.key = key
		c.byKey[key] = item
		c.size++
		c.increment(item)
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
