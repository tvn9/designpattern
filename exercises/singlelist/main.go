package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
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
	node.next = nil

	// check if there no linkedlist exist
	if l.head == nil {
		// assign linked list head and tail node's address to linked list
		l.head = node
		l.tail = node
		// add node address to a linked list lists[] for tracking purpose
		l.SetNodeLocation()
	} else {
		// print this message when the is an existing linked list
		fmt.Println("cannot creat new list, because there is an existing list", l.lists)
	}
}

func (l *LinkedList) SetNodeLocation() {
	node := &Node{}
	if l.head != nil {
		l.lists = nil
		for node = l.head; node != nil; node = node.next {
			l.lists = append(l.lists, node)
		}
	}
}

// AddToHead add a node to the head of the LinkedList
func (l *LinkedList) AddToHead(val int) {
	// Initialize a assign value to a new Node
	node := &Node{}
	node.data = val
	node.next = nil
	if l.head != nil {
		node.next = l.head
	} else {
		fmt.Println("list does not exist.")
	}
	l.head = node
	l.SetNodeLocation()
}

// Push insert the Node to the Head if the linkedlist is not exist, otherwise add to end of the LinkedList
func (l *LinkedList) Push(val int) {
	// Initialize and add Data to the new Node
	node := &Node{}
	node.data = val
	node.next = nil

	// Add the new node to the front of the list
	if l.head != nil {
		l.tail.next = node
	} else {
		l.head = node
	}
	l.tail = node
	l.SetNodeLocation()
}

// NodeWithValue returns existing *Node based on input value
func (l *LinkedList) GetNodePosition(p int) *Node {
	node := &Node{}
	for i := range l.lists {
		if i == p {
			return node
		}
	}
	return node
}

// AddAfter insert a node after a predidefined node in an existing LinkedList
func (l *LinkedList) AddAfter(position, val int) {
	// Initialize and add data to the new Node
	node := &Node{}
	node.data = val
	node.next = nil
	for i, v := range l.lists {
		if i == position {
			node.next = v.next
			v.next = node
		}
	}
	/*
		nodeWithVal := l.NodeWithValue(nodeVal)
		if nodeWithVal != nil {
			node.next = nodeWithVal.next
			nodeWithVal.next = node
		}
	*/
	l.SetNodeLocation()
}

// PrintLinkedList iterates through the LinkedList and print the LinkedList data
func (l *LinkedList) PrintLinkedList() {
	fmt.Println("----- Linked List Info -----")
	fmt.Printf("Head: %p\n", l.head)
	fmt.Printf("Data: %d\n", l.head.data)
	fmt.Printf("Tail: %p\n", l.tail)
	fmt.Printf("Linked List length: %d\n", len(l.lists))
	fmt.Println("Linked List node addresses")
	for i, v := range l.lists {
		fmt.Printf("#%d: Head: %p | Tail: %p | Pointer: %p | NextNode: %p | Data:%d\n", i, l.head, l.tail, v, v.next, v.data)
	}
	fmt.Println("")
}

func (l *LinkedList) IterateList() {
	i := 1
	node := &Node{}
	for node = l.head; node != nil; node = node.next {
		fmt.Printf("#%d: Head: %p | Tail: %p | Pointer: %p | NextNode: %p | Data:%d\n", i, l.head, l.tail, node, node.next, node.data)
		i++
	}
}

// Program starting point
func main() {
	// 1. Create a linked list
	list := LinkedList{}

	fmt.Println("----- Testing Push, whern there is no list exist -----")
	fmt.Println("----- Expected push to create a new list ------")
	list.Push(1111)
	list.PrintLinkedList()

	fmt.Println("------ Create a new List ------")
	list.CreateLinkedList(999)
	list.PrintLinkedList()

	fmt.Println("------ Push 1, 2, 3, 4, 5 ------")
	list.Push(1)
	list.Push(2) // Will not work, there is already a linked list
	list.Push(3) // Will add 2 to the beginig of the list, and push 1 to the nextNode
	list.Push(4)
	list.Push(5)
	list.PrintLinkedList()

	fmt.Println("----- AddAfter(2, 888)")
	list.AddAfter(2, 888)
	list.PrintLinkedList()

	// Add a node to the head
	fmt.Println("------ AddToHead 555, 888, 555 ----- Testing duplicate value")
	list.AddToHead(555)
	list.AddToHead(888)
	list.AddToHead(555)
	list.PrintLinkedList()

	fmt.Println("----- AddAfter 3, 666 -----")
	list.AddAfter(3, 666)
	list.PrintLinkedList()
}

// Create a single list without looking at code example, use only the theory
// and psudo code algorithms
//
// Common operation of a linked list
// 1. Creation of Linked List - Done
// 2. Insertion of Linked List - Done (function AddToHead, ??, ??)
// 3. Traversal of Linked List - Done
// 4. Searching in a Linked List -
// 5. Deletion of a node from a Linked List
// 6. Deletion of Linked List
