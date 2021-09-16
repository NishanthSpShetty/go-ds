package cache

import (
	"container/list"
	"fmt"
)

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

func (f *FrequencyItem) String() string {
	return fmt.Sprintf(" frequncy %d, entries %v ", f.freq, f.entries)
}

func (cI *CacheItem) String() string {
	return fmt.Sprintf("{:key %s, :value %v}", cI.key, cI.value)
}
