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
	for pq.Size() > 0 {
		v := pq.Poll()
		top := pq.Peek()
		assert.Greater(t, v, top)
	}
}
