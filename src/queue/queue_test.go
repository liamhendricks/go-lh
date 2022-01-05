package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	var q Queue = newLHQueue()
	assert.Equal(t, true, q.Empty())
	q.Enqueue(0)
	assert.Equal(t, false, q.Empty())
	q.Enqueue(1)
	q.Enqueue(3)
	val := q.Dequeue()
	assert.Equal(t, 0, val)
	assert.Equal(t, 2, q.Size())
}
