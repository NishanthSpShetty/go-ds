package cache

import "container/list"

type CacheItem struct {
	key             string
	value           interface{}
	frequencyParent *list.Element
}

type FrequencyItem struct {
	entries map[*CacheItem]byte
	freq    int
}

type Cache struct {
	byKey    map[string]*CacheItem
	freqs    *list.List
	capacity int
	size     int
}
