package heap

type Heap interface {
	Insert(int)
	Peek() int
	Poll() int
	Size() int
	Print()
}
