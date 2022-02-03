package linked_list

import "fmt"

type LHLinkedList struct {
	tail *Node
	head *Node
}

func NewLHDoubleLinkedList() *LHLinkedList {
	return &LHLinkedList{}
}

type Node struct {
	Key  int
	next *Node
	prev *Node
}

func (n *Node) PrintNext() *Node {
	fmt.Printf("node: %d", n.Key)
	if n.next == nil {
		return nil
	}
	return n.next.PrintNext()
}

func (n *Node) NextNode(node *Node) *Node {
	if n == node {
		return n
	}

	if n.next == nil {
		return nil
	}

	return n.next.NextNode(node)
}

func (n *Node) NextKey(key int) *Node {
	if n.Key == key {
		return n
	}

	if n.next == nil {
		return nil
	}

	return n.next.NextKey(key)
}

func (ll *LHLinkedList) Empty() bool {
	if ll.tail == nil && ll.head == nil {
		return true
	}

	return false
}

func (ll *LHLinkedList) ValueAt(key int) *Node {
	return ll.tail.NextKey(key)
}

func (ll *LHLinkedList) PushFront(node *Node) {
	if ll.Empty() {
		ll.head = node
		ll.tail = node
	} else {
		tmp := ll.head
		ll.head = node
		tmp.next = ll.head
		ll.head.prev = tmp
		ll.head.next = nil
	}
}

func (ll *LHLinkedList) PushBack(node *Node) {
	if ll.Empty() {
		ll.head = node
		ll.tail = node
	} else {
		tmp := ll.tail
		ll.tail = node
		ll.tail.next = tmp
		tmp.prev = ll.tail
		ll.tail.prev = nil
	}
}

func (ll *LHLinkedList) PopFront() *Node {
	if ll.head == nil {
		return nil
	}
	item := ll.head
	ll.head = item.prev
	ll.head.next = nil

	return item
}

func (ll *LHLinkedList) PopBack() *Node {
	if ll.tail == nil {
		return nil
	}
	item := ll.tail
	ll.tail = item.next
	ll.tail.prev = nil

	return item
}

func (ll *LHLinkedList) Front() *Node {
	return ll.head
}

func (ll *LHLinkedList) Back() *Node {
	return ll.tail
}

func (ll *LHLinkedList) Erase(node *Node) {
	nn := node.next
	pn := node.prev
	if nn != nil {
		nn.prev = pn
	} else {
		ll.head = pn
	}
	if pn != nil {
		pn.next = nn
	} else {
		ll.tail = nn
	}
}

func (ll *LHLinkedList) Print() {
	ll.tail.PrintNext()
}
