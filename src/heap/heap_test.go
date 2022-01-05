package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	var pq Heap = newLHPriorityQueue()
	pq.Insert(5)
	pq.Insert(10)
	pq.Insert(8)
	pq.Insert(1)
	pq.Insert(11)
	pq.Insert(9)
	pq.Insert(7)
	pq.Insert(13)

	// check size
	assert.Equal(t, 8, pq.Size())

	// top element should be 1
	assert.Equal(t, 1, pq.Peek())

	// pop 1
	assert.Equal(t, 1, pq.Poll())

	// top element should be 5
	assert.Equal(t, 5, pq.Peek())

	// should have 7
	assert.Equal(t, 7, pq.Size())
	pq.Print()
}
