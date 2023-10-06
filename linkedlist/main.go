package main

import "fmt"

// Node
type Node struct {
	value int
	next  *Node
}

// LinkedList
type LinkedList struct {
	head  *Node
	tail  *Node
	nodes int
}

// AddToHead
func (l *LinkedList) AddToHead(val int) {
	var newNode = Node{}
	newNode.value = val
	newNode.next = nil

	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
	} else {
		newNode.next = l.head
		l.head = &newNode
	}
	l.nodes++
}

// LastNode
func (l *LinkedList) LastNode() *Node {
	lastNode := l.tail
	return lastNode
}

// HeadNode
func (l *LinkedList) HeadNode() *Node {
	var headNode = l.head
	return headNode
}

// AddAfter
func (l *LinkedList) AddAfter(eVal, val int) {
	node := &Node{}

	node.value = val
	node.next = nil

}

// AddToTail
func (l *LinkedList) AddToTail(val int) {
	node := &Node{}

	node.value = val
	node.next = nil

	if l.head == nil {
		fmt.Println("The list is empty.")
		return
	} else {
		l.tail.next = node
		l.tail = node
	}
	l.nodes++
}

// NodeWithValue
func (l *LinkedList) NodeWithValue(val int) *Node {
	var nodeWithVal *Node

	if l.head == nil {
		fmt.Println("Something went wrong.")
		return nodeWithVal
	}

	for node := l.head; node != nil; node = node.next {
		if node.value == val {
			nodeWithVal = node
		}
	}
	return nodeWithVal
}

// IterateList
func (l *LinkedList) InterateList() {
	var node *Node
	for node = l.head; node != nil; node = node.next {
		fmt.Println("head addr", &l.head, node.value, &node.value, node.next)
	}
}

func main() {

	list := LinkedList{}

	list.AddToHead(1)
	list.AddToHead(2)
	list.AddToTail(9)
	list.AddToTail(5)

	// Get headNode
	hNode := list.HeadNode()
	fmt.Println("Head Node:", hNode)

	// Get lastNode
	lNode := list.LastNode()
	fmt.Println("Last Node:", lNode)

	list.InterateList()

	fmt.Println(list.NodeWithValue(2))
}
