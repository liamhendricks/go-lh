package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapSort(t *testing.T) {
	unsorted := []int{2, 8, 3, 1, 9, 5, 6, 0, 5, 5, 7}
	sorted := heapSort(unsorted)
	assertArraySorted(sorted, t, true)
}

func assertArraySorted(sorted []int, t *testing.T, ascending bool) {
	for i, s := range sorted {
		var n int
		if i+1 >= len(sorted) {
			n = i
		} else {
			n = i + 1
		}
		if ascending {
			assert.GreaterOrEqual(t, s, sorted[n])
		} else {
			assert.LessOrEqual(t, s, sorted[n])
		}
	}
}
