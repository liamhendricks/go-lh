package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	unsorted := []int{8, 15, 7, 28, 3, 98, 12}
	sorted := mergeSort(unsorted)
	assert.NotEqual(t, sorted, unsorted)
}
