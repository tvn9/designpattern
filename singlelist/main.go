// Single LinkedList Implement
package main

import (
	"fmt"
)

// Node data structure
type Node struct {
	value int
	next  *Node
}

// LinkedList data structure
type LinkedList struct {
	head  *Node
	tail  *Node
	nodes int
}

// AddToHead adds the new node to beginning of the list
func (l *LinkedList) AddToHead(val int) {
	var node = Node{}
	node.value = val
	node.next = nil

	if l.head != nil {
		node.next = l.head
	}
	l.head = &node
	l.tail = &node
	l.nodes++
}

// AddToEnd adds the new node with value to the end of the list
func (l *LinkedList) AddToEnd(val int) {
	var lastNode *Node
	node := &Node{}
	node.value = val
	node.next = nil
	lastNode = l.LastNode()
	if lastNode != nil {
		lastNode.next = node
	}
	l.nodes++
}

// NodeWithValue returns Node given parameter value
func (l *LinkedList) NodeWithValue(val int) *Node {
	var node, nodeWithVal *Node

	for node = l.head; node != nil; node = node.next {
		if node.value == val {
			nodeWithVal = node
			break
		}
	}
	return nodeWithVal
}

// AddAfter addes a node with value after a given node in the list
func (l *LinkedList) AddAfter(nodeVal int, val int) {
	var nodeWithVal *Node
	node := &Node{}

	node.value = val
	node.next = nil

	nodeWithVal = l.NodeWithValue(nodeVal)
	if nodeWithVal != nil {
		node.next = nodeWithVal.next
		nodeWithVal.next = node
	}
	l.nodes++
}

// IterateList method iterates over Linkedlist
func (l *LinkedList) IterateList() {
	var node *Node
	for node = l.head; node != nil; node = node.next {
		fmt.Println(&node.value, node.value)
	}
}

// FirstNode returns the first Node in the linkedlist
func (l *LinkedList) HeadNode() *Node {
	return l.head
}

func (n *Node) NextNode() *Node {
	return n.next
}

// LastNode returns the last Node
func (l *LinkedList) LastNode() *Node {
	var node *Node
	for node = l.head; node != nil; node = node.next {
		if node.next == nil {
			l.tail = node
		}
	}
	return l.tail
}

func main() {
	// Create a blank linkedList
	list := LinkedList{}

	list.AddToHead(1)
	list.AddToEnd(2)
	list.AddToEnd(3)
	list.AddToEnd(4)
	list.AddToEnd(9)

	list.IterateList()

	fmt.Println(list.LastNode())
}
