package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	lru := NewLHLRUCache(10)
	for i := 0; i < 25; i++ {
		lru.Set(i)
	}

	// last 10 set items should be in the cache
	for i := 24; i > 14; i-- {
		node := lru.Get(i)
		assert.NotNil(t, node)
	}

	lru.Set(44)
	node := lru.Get(44)
	assert.NotNil(t, node)

	// 24 should have popped off
	node = lru.Get(24)
	assert.Nil(t, node)
}
