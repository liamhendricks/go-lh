package bst

import (
	"fmt"
	"sync"

	"github.com/liamhendricks/go-lh/src/queue"
)

// threadsafe BST
type Bst interface {
	Insert(int, string)
	Remove(int) *Node
	Search(int) *Node
	Find(string) *Node
	Min() *Node
	Max() *Node
	InOrderTraverse()
	PreOrderTraverse()
	PostOrderTraverse()
	Print()
}

type LHBst struct {
	root *Node
	lock sync.RWMutex
}

type Node struct {
	key   int
	val   string
	left  *Node
	right *Node
}

func newLHBst() *LHBst {
	return &LHBst{
		root: nil,
		lock: sync.RWMutex{},
	}
}

// insert item in tree at key
func (bst *LHBst) Insert(key int, val string) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	n := &Node{
		key: key,
		val: val,
	}

	if bst.root == nil {
		bst.root = n
	} else {
		insertNode(bst.root, n)
	}
}

func (bst *LHBst) Remove(key int) *Node {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	return remove(bst.root, key)
}

// remove a node. i believe worst case this is O(nlogn)
func remove(node *Node, key int) *Node {
	if node == nil {
		return nil
	}

	if key < node.key {
		node.left = remove(node.left, key)
		return node
	}

	if key > node.key {
		node.right = remove(node.right, key)
		return node
	}

	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}

	if node.left == nil {
		node = node.right
		return node
	}

	if node.right == nil {
		node = node.left
		return node
	}

	leftmostrightside := node.right
	for {
		//find smallest value on the right side
		if leftmostrightside != nil && leftmostrightside.left != nil {
			leftmostrightside = leftmostrightside.left
		} else {
			break
		}
	}

	node.key, node.val = leftmostrightside.key, leftmostrightside.val
	node.right = remove(node.right, node.key)
	return node
}

// search by key using binary search
func (bst *LHBst) Search(key int) *Node {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	return search(bst.root, key)
}

// find item by value, using bfs
func (bst *LHBst) Find(item string) *Node {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	fmt.Printf("looking for: %s\n", item)
	return breadthFirst(bst.root, item)
}

func (bst *LHBst) Max() *Node {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	return goRight(bst.root)
}

func (bst *LHBst) Min() *Node {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	return goLeft(bst.root)
}

func (bst *LHBst) InOrderTraverse() {
	inOrder(bst.root)
}

func (bst *LHBst) PreOrderTraverse() {
	preOrder(bst.root)
}

func (bst *LHBst) PostOrderTraverse() {
	postOrder(bst.root)
}

// bfs O(n)
func breadthFirst(root *Node, item string) *Node {
	visited := make(map[int]*Node)
	q := &queue.LHQueue{}
	q.Enqueue(root)
	visited[root.key] = root

	for {
		if q.Empty() {
			break
		}
		node := q.Dequeue().(*Node)
		fmt.Printf("searching: %s\n", node.val)
		if node.val == item {
			fmt.Printf("found: %s\n", item)
			return node
		}

		if node.left != nil {
			if _, ok := visited[node.left.key]; !ok {
				visited[node.left.key] = node.left
				q.Enqueue(node.left)
			}
		}

		if node.right != nil {
			if _, ok := visited[node.right.key]; !ok {
				visited[node.right.key] = node.right
				q.Enqueue(node.right)
			}
		}
	}

	fmt.Printf("not found: %s\n", item)
	return nil
}

// in order dfs O(n)
func inOrder(node *Node) {
	if node != nil {
		inOrder(node.left)
		fmt.Printf("key: %d\n", node.key)
		inOrder(node.right)
	}
}

// pre order dfs O(n)
func preOrder(node *Node) {
	if node != nil {
		fmt.Printf("key: %d\n", node.key)
		preOrder(node.left)
		preOrder(node.right)
	}
}

// post order dfs O(n)
func postOrder(node *Node) {
	if node != nil {
		postOrder(node.left)
		postOrder(node.right)
		fmt.Printf("key: %d\n", node.key)
	}
}

// binary insert O(logn)
func insertNode(node, newNode *Node) {
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

// binary search O(logn)
func search(node *Node, key int) *Node {
	if node == nil {
		return nil
	}

	if node.key > key {
		return search(node.left, key)
	}

	if node.key < key {
		return search(node.right, key)
	}

	return node
}

func goRight(node *Node) *Node {
	if node.right == nil {
		return node
	} else {
		return goRight(node.right)
	}
}

func goLeft(node *Node) *Node {
	if node.left == nil {
		return node
	} else {
		return goLeft(node.left)
	}
}

// String prints a visual representation of the tree using 3 types of dfs
func (bst *LHBst) Print() {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	fmt.Println("IN ORDER")
	inOrder(bst.root)
	fmt.Println("------------------------------------------------")
	fmt.Println("PRE ORDER")
	preOrder(bst.root)
	fmt.Println("------------------------------------------------")
	fmt.Println("POST ORDER")
	postOrder(bst.root)
	fmt.Println("------------------------------------------------")
}
