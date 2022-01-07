package sorting

import (
	"github.com/liamhendricks/go-lh/src/heap"
)

// Altogether O(nlogn)
func heapSort(nums []int) []int {
	var sorted []int
	// priority queue for sorting
	pq := buildMaxHeap(nums)
	for pq.Size() > 0 {
		// pq.Poll() pops our highest priority item, and runs heapifyDown()
		// pq.Poll() runs O(logn), but we call n-1 times
		sorted = append(sorted, pq.Poll())
	}
	return sorted
}

// O(n)
func buildMaxHeap(nums []int) heap.Heap {
	pq := heap.NewLHPriorityQueue()
	for _, n := range nums {
		pq.Insert(n)
	}

	return pq
}
