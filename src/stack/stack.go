package stack

import (
	"fmt"
)

type Stack interface {
	Push(interface{})
	Pop() interface{}
	Top() interface{}
	Empty() bool
	Print()
}

type LHStack struct {
	name  string
	items []interface{}
}

// append onto stack
func (l *LHStack) Push(data interface{}) {
	l.items = append(l.items, data)
}

// remove most recently added item and return it
func (l *LHStack) Pop() interface{} {
	if !l.Empty() {
		item := l.items[len(l.items)-1]
		l.items = l.items[:len(l.items)-1]
		return item
	}

	return nil
}

func (l *LHStack) Top() interface{} {
	if !l.Empty() {
		return l.items[len(l.items)-1]
	}

	return nil
}

func (l *LHStack) Empty() bool {
	return len(l.items) == 0
}

func (l *LHStack) Print() {
	fmt.Printf("Printing %s\n", l.name)
	for _, item := range l.items {
		fmt.Printf("item: %s\n", item.(string))
	}
}

func newLHStack() *LHStack {
	var items []interface{}
	return &LHStack{
		name:  "LHStack",
		items: items,
	}
}
