package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_cache_set(t *testing.T) {
	c := New(100)
	c.Set("a", 101)

	assert.Equal(t, c.Get("a"), 101, "can set and get the value associated with key")
	c.Set("a", 200)
	assert.Equal(t, c.Get("a"), 200, "can update and get the value associated with key")
}
