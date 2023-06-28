package main

import (
	"fmt"
)

type Node struct {
	data     int
	nextNode *Node
}

type LinkedList struct {
	head  *Node
	tail  *Node
	lists []*Node
}

// CreateLinkedList only create a new Linked List, an prompt message will print if there
// is an existing Linked List
func (l *LinkedList) CreateLinkedList(val int) {
	// Initialize and add data to a the new node
	node := &Node{}
	node.data = val
	node.nextNode = nil

	// check if there no linkedlist exist
	if len(l.lists) != 1 {
		// assign linked list head and tail node's address to linked list
		l.head = node
		l.tail = node
		// add node address to a linked list lists[] for tracking purpose
		l.lists = append(l.lists, node)
	} else {
		// print this message when the is an existing linked list
		fmt.Println("cannot creat new list, because there is an existing list", l.lists)
	}
}

func (l *LinkedList) AddToHead(val int) {
	// Initialize and Add Data to the new Node
	node := &Node{}
	fmt.Printf("Node address %p\n", node)
	node.data = val
	node.nextNode = nil

	// Add the new node to the front of the list
	if l.head != nil {
		node.nextNode = l.head
	}
	l.head = node
	l.lists = append(l.lists, node)
}

func (l *LinkedList) PrintLinkedList() {
	fmt.Println("----- Linked List Info -----")
	fmt.Printf("Head: %p\n", l.head)
	fmt.Printf("Data: %d\n", l.head.data)
	fmt.Printf("Tail: %p\n", l.tail)
	fmt.Printf("Linked List length: %d\n", len(l.lists))
	fmt.Println("Linked List node addresses")
	for i, v := range l.lists {
		fmt.Printf("Node #%d - Pointer: %p - Data: %d\n", i, v, v.data)
	}
}

func main() {
	// 1. Create a linked list
	list := LinkedList{}

	list.CreateLinkedList(1)
	list.CreateLinkedList(2)
	list.AddToHead(2)
	list.AddToHead(3)
	list.AddToHead(4)

	list.PrintLinkedList()

}

// Create a single list without looking at code example, use only the theory
// and psudo code algorithms
//
// Common operation of a linked list
// 1. Creation of Linked List - Done
// 2. Insertion of Linked List
// 3. Traversal of Linked List
// 4. Searching in a Linked List
// 5. Deletion of a node from a Linked List
// 6. Deletion of Linked List
