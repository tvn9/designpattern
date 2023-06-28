// Double LinkedList
package main

// Node struct
type Node struct {
	number   int
	nextNode *Node
	preNode  *Node
}

// LinkedList struct
type LinkedList struct {
	headNode *Node
}

//Create a link list
func (l *LinkedList) CreateLinkedList(nodeData int) {
	// initialize a node
	
}

// NodeBetweenValues mothod of LinkedList
func (l *LinkedList) NodeBetweenValues(firstNumber int, secondNumber int) *Node {
	var node, nodeWith *Node
	for node = l.headNode; node != nil; node = node.nextNode {
		if node.preNode != nil && node.nextNode != nil {
			if node.preNode.number == firstNumber && node.nextNode.number == secondNumber {
				nodeWith = node
				break
			}
		}
	}
	return nodeWith
}

// AddToHead add the node the head of the list
func (l *LinkedList) AddToHead(num int) {
	var node = &Node{}
	node.number = num
	node.nextNode = nil

	if l.headNode != nil {
		node.nextNode = l.headNode
		l.headNode.preNode = node
	}
	l.headNode = node
}

// NodeWithValue
func (l *LinkedList) NodeWithValue(nodeVal int) *Node {
	var node = &Node{}
	var nodeWithVal *Node
	for node = l.headNode; node != nil; node = node.nextNode {
		if node.number = nodeVal {
			nodeWithVal = node
			break
		}
	}
	return nodeWithVal
}

// AddAfter
func (l *LinkedList) AddAfter(nodeVal int, newVal int) {
	var (
		node = &Node{}
		nodeWithVal *Node
	)
	node.number = newVal
	node.nextNode = nil

	nodeWithVal = l.NodeWithValue(nodeVal)
	if nodeWithValue != nil {
		node.previousNode = nodeWithVal
		nodeWithVal.nextNode = node
	}
}

// HeadNode returns the pointer to the first node 
func (l *LinkedList) HeadNode() *Node {
	return l.headNode
}

// NextNode returns the pointer to the next node in the list
func (l *LinkedList) NextNode() *Node {
	return l.headNode.nextNode
}


// main method
func main() {
	list := LinkedList{}
	list.AddToHead(1)


	// get the head node
	hn := l.HeadNode()

	// get the next node
	nn := l.NextNode()

	fmt.Printf("Head Node address %p, data: %d\n", hn, hn.number)
}
