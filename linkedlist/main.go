package main

import "fmt"

// Node struct holds the node value and next note pointer address
type Node struct {
	value int
	next  *Node
}

// LinkedList holds variables point the first node and last node
type LinkedList struct {
	head  *Node
	tail  *Node
	nodes int
}

// AddToHead adds the new node to the begining of the list
func (l *LinkedList) AddToHead(val int) {
	node := &Node{
		value: val,
		next:  nil,
	}
	if l.head != nil {
		node.next = l.head
	}
	l.head = node
	l.tail = node
	l.nodes++
}

// HeadNode return the head node
func (l *LinkedList) HeadNode() *Node {
	var headNode = l.head
	return headNode
}

// LastNode return the last node of the list
func (l *LinkedList) LastNode() *Node {
	var lastNode *Node
	if l.head != nil {
		for node := l.head; node != nil; node = node.next {
			if node.next == nil {
				lastNode = node
			}
		}
	}
	return lastNode
}

// AddAfter adds the new node to the end of the list
func (l *LinkedList) AddAfter(eVal int, val int) {
	var nodeWithValue *Node
	node := &Node{
		value: val,
		next:  nil,
	}
	nodeWithValue = l.NodeWithValue(eVal)
	if nodeWithValue != nil {
		node.next = nodeWithValue.next
		nodeWithValue.next = node
	}
	l.nodes++
}

// AddToTail adds new node to end of the list
func (l *LinkedList) AddToEnd(val int) {
	node := &Node{
		value: val,
		next:  nil,
	}
	lastNode := l.LastNode()
	lastNode.next = node
	l.nodes++
}

// NodeWithValue return the node with the value
func (l *LinkedList) NodeWithValue(nVal int) *Node {
	var node, nodeWithVal *Node
	if l.head != nil {
		for node = l.head; node != nil; node = node.next {
			if node.value == nVal {
				nodeWithVal = node
				break
			}
		}
	}
	return nodeWithVal
}

// UpdateNode
func (l *LinkedList) UpdateNode(oVal, nVal int) {
	nodeWithVal := l.NodeWithValue(oVal)
	if nodeWithVal.value != oVal {
		fmt.Println("Can not find node with value:", oVal)
	} else {
		nodeWithVal.value = nVal
	}
}

// IterateList
func (l *LinkedList) InterateList() {
	var node *Node
	for node = l.head; node != nil; node = node.next {
		fmt.Println("head addr", &l.head, node, node.next)
	}
}

func main() {

	list := &LinkedList{}

	list.AddToHead(2)
	list.AddToHead(1)
	list.AddToEnd(9)
	list.AddAfter(2, 3)
	list.AddAfter(1, 4)
	list.AddAfter(3, 5)
	list.AddToEnd(100)
	list.AddToHead(10)

	list.UpdateNode(10, 10)
	list.UpdateNode(1, 20)
	list.UpdateNode(4, 30)
	list.UpdateNode(2, 40)
	list.UpdateNode(3, 50)
	list.UpdateNode(5, 60)
	list.UpdateNode(9, 70)
	list.UpdateNode(100, 80)

	list.InterateList()

	// Get headNode
	hNode := list.HeadNode()
	fmt.Println("Head Node:", hNode)

	// Get lastNode
	lNode := list.LastNode()
	fmt.Println("Last Node:", lNode)

	fmt.Println(list.NodeWithValue(10))
	fmt.Println(list.nodes)
}
