package bst

import (
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
)

func TestBst(t *testing.T) {
	var lhBst Bst = newLHBst()

	var startingValues = []int{10, 5, 15, 8, 13, 14, 12, 20, 18, 25, 3, 0, 4}
	for _, v := range startingValues {
		lhBst.Insert(v, fake.FullName())
	}

	max := lhBst.Max()
	assert.Equal(t, 25, max.key)

	lhBst.Insert(26, fake.FullName())

	max = lhBst.Max()
	assert.Equal(t, 26, max.key)

	n := lhBst.Find("eeeeeeoooooooooooooeeeeeeeeeeeooooooooo")
	assert.Nil(t, n)

	found := lhBst.Search(20)
	assert.NotNil(t, found)
	notFound := lhBst.Search(100)
	assert.Nil(t, notFound)

	lookFor := lhBst.Find(found.val)
	assert.NotNil(t, lookFor)
	assert.Equal(t, found.val, lookFor.val)

	// remove leaf
	lhBst.Remove(0)
	found = lhBst.Search(0)
	assert.Nil(t, found)

	// remove node with 1 child
	lhBst.Remove(25)
	found = lhBst.Search(25)
	assert.Nil(t, found)

	// remove node with 2 children
	lhBst.Remove(13)
	found = lhBst.Search(13)
	assert.Nil(t, found)

	lhBst.Print()
}
