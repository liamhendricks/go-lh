package heap

import (
	"fmt"
	"sync"
)

// priority queue using a max heap
type LHPriorityQueue struct {
	data []int
	lock sync.RWMutex
}

func NewLHPriorityQueue() *LHPriorityQueue {
	return &LHPriorityQueue{
		data: []int{},
		lock: sync.RWMutex{},
	}
}

func (pq *LHPriorityQueue) Size() int {
	return len(pq.data)
}

// Insert new element and heapify up, O(logn)
func (pq *LHPriorityQueue) Insert(key int) {
	pq.lock.Lock()
	defer pq.lock.Unlock()

	pq.insert(key)
}

func (pq *LHPriorityQueue) insert(key int) {
	pq.data = append(pq.data, key)
	pq.heapifyUp()
}

func (pq *LHPriorityQueue) swap(i, j int) {
	tmp := pq.data[i]
	pq.data[i] = pq.data[j]
	pq.data[j] = tmp
}

// Return top element, O(1)
func (pq *LHPriorityQueue) Peek() int {
	pq.lock.Lock()
	defer pq.lock.Unlock()

	if len(pq.data) == 0 {
		return -1
	}

	return pq.data[0]
}

// Remove root element and heapify down, O(logn)
func (pq *LHPriorityQueue) Poll() int {
	pq.lock.Lock()
	defer pq.lock.Unlock()

	if len(pq.data) == 0 {
		return -1
	}

	item := pq.data[0]
	pq.data[0] = pq.data[len(pq.data)-1]
	pq.data = pq.data[:len(pq.data)-1]
	pq.heapifyDown()

	return item
}

func (pq *LHPriorityQueue) heapifyDown() {
	i := 0
	for leftChildIndex(i) < pq.Size() {
		var k int = rightChildIndex(i)
		var j int = leftChildIndex(i)
		if k < pq.Size() && pq.data[j] < pq.data[k] {
			j++
		}
		if pq.data[i] > pq.data[j] {
			break
		}
		pq.swap(i, j)
		i = j
	}
}

func (pq *LHPriorityQueue) heapifyUp() {
	i := len(pq.data) - 1

	for i > 0 && pq.parent(i) < pq.data[i] {
		pq.swap(parentIndex(i), i)
		i = parentIndex(i)
	}
}

func parentIndex(child int) int {
	return (child - 1) / 2
}

func leftChildIndex(parent int) int {
	return 2*parent + 1
}

func rightChildIndex(parent int) int {
	return 2*parent + 2
}

func (pq *LHPriorityQueue) leftChild(i int) int {
	return pq.data[leftChildIndex(i)]
}

func (pq *LHPriorityQueue) rightChild(i int) int {
	return pq.data[rightChildIndex(i)]
}

func (pq *LHPriorityQueue) parent(i int) int {
	return pq.data[parentIndex(i)]
}

func (pq *LHPriorityQueue) Print() {
	for k, v := range pq.data {
		fmt.Printf("%d: %d\n", k, v)
	}
}
