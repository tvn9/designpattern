// Single LinkedList Implement
package main

import (
	"fmt"
)

type Node struct {
	number   int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
	nodes    int
}

func (l *LinkedList) AddToHead(val int) {
	var node = Node{}
	node.number = val
	node.nextNode = nil

	if l.headNode != nil {
		node.nextNode = l.headNode
	}
	l.headNode = &node
	l.nodes++
}

// AddToEnd adds the node with value to the end of the list
func (l *LinkedList) AddToEnd(val int) {
	var lastNode *Node
	node := &Node{}
	node.number = val
	node.nextNode = nil
	lastNode = l.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	}
	l.nodes++
}

// NodeWithValue returns Node given parameter value
func (l *LinkedList) NodeWithValue(nval int) *Node {
	var node, nodeWith *Node

	for node = l.headNode; node != nil; node = node.nextNode {
		if node.number == nval {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

// AddAfter addes a node with value after a given node in the list
func (l *LinkedList) AddAfter(nval int, val int) {
	var nodeWith *Node
	node := &Node{}

	node.number = val
	node.nextNode = nil

	nodeWith = l.NodeWithValue(nval)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
	l.nodes++
}

// IterateList method iterates over Linkedlist
func (l *LinkedList) IterateList() {
	var node *Node
	for node = l.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.number)
	}
}

// FirstNode returns the first Node in the linkedlist
func (l *LinkedList) HeadNode() *Node {
	return l.headNode
}

func (n *Node) NextNode() *Node {
	return n.nextNode
}

// LastNode returns the last Node
func (l *LinkedList) LastNode() *Node {
	var node, lastNode *Node
	for node = l.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

func (l *LinkedList) Insert(data int) {

}

func main() {
	// Create a blank linkedList
	list := LinkedList{}

	list.AddToHead(1)
	list.AddToHead(2)

	list.AddAfter(1, 3)
	list.AddAfter(3, 4)
	list.AddToEnd(5)
	list.AddToEnd(6)

	h := list.HeadNode()
	n := list.headNode.NextNode()

	fmt.Println("First", h.number)
	fmt.Println("Next", n.number)
	fmt.Println("Number of nodes:", list.nodes)

	list.IterateList()

	lastNode := list.LastNode()

	fmt.Println("Last Node ", lastNode.number)
}
