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

func TestStaticQueue(t *testing.T) {
	var q Queue = NewStaticQueue(10)
	assert.Equal(t, true, q.Empty())
	for i := 0; i < 10; i++ {
		q.Enqueue(int64(i))
	}
	assert.Equal(t, false, q.Empty())
	assert.Equal(t, q.Size(), 10)
	q.Enqueue(int64(10))
	// size should remain 10, popping off the last element
	assert.Equal(t, q.Size(), 10)
	val := q.Dequeue()
	// popped element should not be 0, since it should have auto popped
	assert.Equal(t, val, int64(1))

	q = NewStaticQueue(10)
	for i := 0; i < 10; i++ {
		q.Enqueue(int64(i))
	}
	a := q.Average()
	assert.Equal(t, a, float32(4.5))
}
