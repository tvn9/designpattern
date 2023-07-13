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

// Insert inserts the Node to the Head if the linkedlist is not exist, otherwise add to end of the LinkedList
func (l *LinkedList) Insert(val int) {
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
	l.SetNodeLocation()
	fmt.Println("")
}

func (l *LinkedList) AddToLast(val int) {
	// Initialize a new Node
	node := &Node{}
	node.data = val
	node.next = nil

	if l.tail != nil {
		l.tail.next = node
	}
	l.tail = node
	l.SetNodeLocation()
}

func (l *LinkedList) DeleteNode(val int, position int) {
	if l.head != nil {
		// Find the node targeted for delete
		for i, v := range l.lists {
			// Compare the search parameter with current node values and position
			if position == i && val == v.data {
				// remove the node by re-assign the linkedlist's head and next pointers
				l.lists[i-1].next = v.next
			}
		}
		l.SetNodeLocation()
	}
}

// IterateList cycle through the LinkedList and print out list info
func (l *LinkedList) IterateList() {
	i := 1
	var node *Node
	for node = l.head; node != nil; node = node.next {
		fmt.Printf("#%d: Head: %p | Tail: %p | Pointer: %p | NextNode: %p | Data:%d\n", i, l.head, l.tail, node, node.next, node.data)
		i++
	}
}

// SearchNode
func (l *LinkedList) SearchNode(val int) bool {
	var node *Node
	var result bool
	if l.head != nil {
		for node = l.head; node != nil; node = node.next {
			if node.data == val {
				// fmt.Printf("Search %d, Found %d, Address %p\n", val, node.data, node)
				result = true
				break
			}
		}
	}
	return result
}

// FistNode returns the first node
func (l *LinkedList) FirstNode() *Node {
	return l.head
}

// NextNode returns the next node
func (l *LinkedList) NextNode() *Node {
	return l.head.next
}

// LastNode returns the last node
func (l *LinkedList) LastNode() *Node {
	return l.tail
}

// Program starting point
func main() {
	// 1. Create a linked list
	list := LinkedList{}

	fmt.Println("----- Testing Push, whern there is no list exist -----")
	fmt.Println("----- Expected push to create a new list ------")
	list.Insert(1111)
	list.PrintLinkedList()

	fmt.Println("------ Create a new List ------")
	list.CreateLinkedList(999)
	list.PrintLinkedList()

	fmt.Println("------ Push 1, 2, 3, 4, 5 ------")
	list.Insert(1)
	list.Insert(2) // Will not work, there is already a linked list
	list.Insert(3) // Will add 2 to the beginig of the list, and push 1 to the nextNode
	list.Insert(4)
	list.Insert(5)
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

	fmt.Println("----- AddToLast -----")
	list.AddToLast(10)
	list.PrintLinkedList()

	fmt.Println("----- FistNode -----")
	fn := list.FirstNode()
	fmt.Printf("First Node %p | Data: %d\n", fn, fn.data)

	fmt.Println("----- NextNode -----")
	nn := list.NextNode()
	fmt.Printf("Next Node %p | Data: %d\n", nn, nn.data)

	fmt.Println("----- LastNode -----")
	ln := list.LastNode()
	fmt.Printf("Last Node: %p | Data: %d\n", ln, ln.data)

	num := 555
	fmt.Println("----- SearchNode -----")
	n := list.SearchNode(num)
	fmt.Println("Item found =", n)

	nodes := []int{888, 555, 1111, 666, 1, 2, 3, 4, 5, 10, 555}
	fmt.Println("----- Search multiple nodes -----")
	for _, n := range nodes {
		result := list.SearchNode(n)
		if result {
			fmt.Println("Item found", n)
		} else {
			fmt.Println("Item not found")
		}
	}
	list.IterateList()

	// Delete a node
	list.DeleteNode(666, 4)
	list.PrintLinkedList()

	list.DeleteNode(888, 6)
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
