// Linked List
package main

import (
	"fmt"
)

// Node class
type Node struct {
	property int
	nextNode *Node
}

// Linkedlist class
type LinkedList struct {
	headNode *Node
}

// AddtoHead method of LinkedList class
func (ll *LinkedList) AddToHead(property int) {
	var node = Node{}
	node.property = property
	if node.nextNode != nil {
		node.nextNode = ll.headNode
	}
	ll.headNode = &node
}

// IterateList method interates over LinkedList
func (ll *LinkedList) IterateList() {
	var node *Node
	for node = ll.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

// LastNode method returns the last Node
func (ll *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = ll.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// AddToEnd method adds the node with property to the end
func (ll *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var lastNode *Node
	lastNode = ll.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	}
}

// NodeWithValue method returns Node given parameter property
func (ll *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWith *Node
	for node = ll.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

// AddAfter method adds a node with nodeProperty after node with property
func (ll *LinkedList) AddAfter(nodeProperty int, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var nodeWith *Node
	nodeWith = ll.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
}

// main method
func main() {
	var ll = LinkedList{}

	ll.AddToHead(1)
	ll.AddToHead(3)
	ll.AddToEnd(5)
	ll.AddAfter(1, 7)
	ll.IterateList()
	fmt.Println(ll.headNode.property)
}
