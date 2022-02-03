package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	ll := NewLHDoubleLinkedList()
	nodes := []*Node{{Key: 5}, {Key: 18}, {Key: 3}, {Key: 28}}

	for _, node := range nodes {
		ll.PushFront(node)
	}

	assert.Equal(t, false, ll.Empty())
	front := ll.Front()
	assert.Equal(t, 28, front.Key)
	back := ll.Back()
	assert.Equal(t, 5, back.Key)

	fp := ll.PopFront()
	assert.Equal(t, 28, fp.Key)

	front = ll.Front()
	assert.Equal(t, 3, front.Key)

	bp := ll.PopBack()
	assert.Equal(t, 5, bp.Key)

	back = ll.Back()
	assert.Equal(t, 18, back.Key)

	ll.Erase(nodes[1])

	back = ll.Back()
	assert.Equal(t, 3, back.Key)

	front = ll.Front()
	assert.Equal(t, 3, front.Key)
}
