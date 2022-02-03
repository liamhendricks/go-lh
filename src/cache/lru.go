package cache

import "github.com/liamhendricks/go-lh/src/linked_list"

type LHLRU struct {
	list     *linked_list.LHLinkedList
	cache    map[int]*linked_list.Node
	capacity int
}

func NewLHLRUCache(capacity int) LHLRU {
	return LHLRU{
		list:     linked_list.NewLHDoubleLinkedList(),
		cache:    make(map[int]*linked_list.Node),
		capacity: capacity,
	}
}

func (c *LHLRU) Get(key int) *linked_list.Node {
	item, ok := c.cache[key]
	if ok {
		front := c.list.Front()
		if front != item {
			c.list.Erase(item)
			c.list.PushFront(item)
		}

		return item
	}

	return nil
}

func (c *LHLRU) Set(key int) {
	item, ok := c.cache[key]
	if ok {
		front := c.list.Front()
		if front != item {
			c.list.Erase(item)
			c.list.PushFront(item)
		}
	} else {
		// if we are at cap, pop our LRU node
		if len(c.cache) >= c.capacity {
			popped := c.list.PopBack()
			delete(c.cache, popped.Key)
		}
		node := &linked_list.Node{Key: key}
		c.cache[key] = node
		c.list.PushFront(node)
	}
}
