package queue

type Queue interface {
	Enqueue(interface{})
	Dequeue() interface{}
	Size() int
	Empty() bool
}

type LHQueue struct {
	items []interface{}
}

// append onto the queue
func (queue *LHQueue) Enqueue(item interface{}) {
	queue.items = append(queue.items, item)
}

// remove first item and return it
func (queue *LHQueue) Dequeue() interface{} {
	if !queue.Empty() {
		item := queue.items[0]
		queue.items = queue.items[1:]
		return item
	}

	return nil
}

func (queue *LHQueue) Empty() bool {
	return queue.Size() == 0
}

func (queue *LHQueue) Size() int {
	return len(queue.items)
}

func newLHQueue() *LHQueue {
	var items []interface{}
	return &LHQueue{
		items: items,
	}
}
