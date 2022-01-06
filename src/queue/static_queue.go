package queue

// this queue stays <= maxSize
type StaticQueue struct {
	values  []int64
	maxSize int
}

func NewStaticQueue(length int) *StaticQueue {
	return &StaticQueue{
		values:  []int64{},
		maxSize: length,
	}
}

func (q *StaticQueue) Size() int {
	return len(q.values)
}

func (q *StaticQueue) Empty() bool {
	return q.Size() == 0
}

func (q *StaticQueue) max() int {
	return q.maxSize
}

func (q *StaticQueue) Dequeue() interface{} {
	if !q.Empty() {
		val := q.values[0]
		q.values = q.values[1:]
		return val
	}

	return -1
}

func (q *StaticQueue) Enqueue(n interface{}) {
	if q.Size() == q.max() {
		_ = q.Dequeue()
	}
	val, ok := n.(int64)
	if ok {
		q.values = append(q.values, val)
	}
}

func (q *StaticQueue) Average() float32 {
	var avg float32
	for _, n := range q.values {
		avg += float32(n)
	}
	return avg / float32(q.Size())
}
