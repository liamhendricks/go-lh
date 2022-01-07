package sorting

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	unsorted := []int{8, 15, 7, 28, 3, 98, 12}
	sorted := mergeSort(unsorted)
	assertArraySorted(sorted, t, false)
}
