package cache

import (
	"container/list"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func printFrequncyList(l *list.List) {
	fmt.Println("Frequency list")

	for front := l.Front(); ; front = front.Next() {
		if front == nil {
			break
		}
		fmt.Printf(": %v\n", front.Value)
	}
}

func Test_cache_set(t *testing.T) {
	c := New(100)
	c.Set("a", 101)

	assert.Equal(t, c.Get("a"), 101, "can set and get the value associated with key")
	printFrequncyList(c.freqs)
	c.Set("a", 200)
	assert.Equal(t, c.Get("a"), 200, "can update and get the value associated with key")

	c.Set("b", 200)
	c.Set("c", 200)
	printFrequncyList(c.freqs)
	c.Set("b", 200)
	printFrequncyList(c.freqs)
}
