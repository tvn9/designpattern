// Doubly LinkedList
package main

import "fmt"

// Node struct
type Node struct {
	data     int
	next     *Node
	previous *Node
}

// LinkedList struct
type LinkedList struct {
	head *Node
	tail *Node
}

// NodeBetweenValues mothod of LinkedList
func (l *LinkedList) NodeBetweenValues(num1 int, num2 int) *Node {
	var node, nodeWith *Node
	for node = l.head; node != nil; node = node.next {
		if node.previous != nil && node.next != nil {
			if node.previous.data == num1 && node.next.data == num2 {
				nodeWith = node
			}
		}
	}
	return nodeWith
}

// AddToHead add the node the head of the list
func (l *LinkedList) AddToHead(num int) {
	node := &Node{}
	node.data = num
	node.next = nil

	if l.head != nil {
		node.next = l.head
		l.head.previous = node
		l.tail = l.head.next
	}
	l.head = node
}

// NodeWithValue returns a Node with given value
func (l *LinkedList) NodeWithValue(val int) *Node {
	var node, nodeWith *Node
	for node = l.head; node != nil; node = node.next {
		if node.data == val {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

// AddAfter inserts a node after a specific node
func (l *LinkedList) AddAfter(nodeVal int, newVal int) {
	node := &Node{}
	var nodeWith *Node

	node.data = newVal
	node.next = nil

	nodeWith = l.NodeWithValue(nodeVal)
	if nodeWith != nil {
		node.previous = nodeWith
		nodeWith.next = node
	}
}

// AddToEnd insert a node to the end of the list
func (l *LinkedList) AddToEnd(val int) {
	node := &Node{}

	node.data = val
	node.next = nil

	lastNode := l.LastNode()

	if lastNode != nil {
		lastNode.next = node
		node.previous = lastNode
	}
	l.tail = node
}

// HeadNode returns the pointer to the first node
func (l *LinkedList) HeadNode() *Node {
	return l.head
}

// NextNode returns the pointer to the next node in the list
func (l *LinkedList) NextNode() *Node {
	return l.head.next
}

// LastNode return the last node from the list
func (l *LinkedList) LastNode() *Node {
	var node *Node
	if l.head != nil {
		for node = l.head; node != nil; node = node.next {
			if node.next == nil {
				l.tail = node
			}
		}
	}
	return l.tail
}

// InterateNodes interates the list and print the node information to terminal
func (l *LinkedList) InterateNodes() {
	if l.head != nil {
		for node := l.head; node != nil; node = node.next {
			fmt.Printf("Node Address %p | Previous %p | Data %d | Next Node %p\n",
				node, node.previous, node.data, node.next)
		}
	}
}

// main method
func main() {
	list := LinkedList{}
	list.AddToHead(1)
	list.AddToHead(2)
	list.AddToHead(3)
	list.AddToHead(4)
	list.AddToHead(5)
	list.AddToEnd(99)

	// get the head node
	hn := list.HeadNode()

	// get the next node
	nn := list.NextNode()

	// get the last node
	ln := list.LastNode()

	fmt.Printf("Head Node address %p, data: %d\n", hn, hn.data)
	fmt.Printf("Next Node address %p, data: %d\n", nn, nn.data)
	fmt.Printf("Last Node address %p, data: %d\n", ln, ln.data)

	// Print the list info
	list.InterateNodes()
}
