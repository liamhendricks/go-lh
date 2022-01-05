package linked_list

type LinkedList interface {
	Empty() bool
	ValueAt(int) interface{}
	PushFront(interface{})
	PushBack(interface{})
	PopFront() interface{}
	PopBack() interface{}
	Front() interface{}
	Back() interface{}
	Insert(int, interface{})
	Erase(int)
	Reverse()
	Remove(interface{})
}

type LHLinkedList struct {
	tail *Node
	head *Node
}

type Node struct {
	key  int
	val  string
	next *Node
}

func newLHDoubleLinkedList() *LHLinkedList {
	return &LHLinkedList{}
}
