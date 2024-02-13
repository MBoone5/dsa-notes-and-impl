
// ------------ Node

type Node struct {
	prev *Node
	next *Node

	Value interface{}
}

func (n *Node) Next() *Node {
	if n == nil {
		return nil
	}

	return n.next
}

func (n *Node) Prev() *Node {
	if n == nil {
		return nil
	}

	return n.prev
}

// ------------ List

type List struct {
	head *Node
	tail *Node
}

func (l *List) Unshift(v interface{}) {
	// Create the new node
	newNode := Node{Value: v, prev: nil, next: l.head}

	if l.head != nil {
		// Update the existing contents
		existingHead := l.head
		existingHead.prev = &newNode
	}

	// Update the list itself
	l.head = &newNode

	if l.tail == nil { // If this is the list's first node
		l.tail = &newNode
	}
}

func (l *List) Push(v interface{}) {
	// Create the new node
	newNode := Node{Value: v, prev: l.tail, next: nil}

	if l.tail != nil {
		// Update the existing contents of the List
		existingTail := l.tail
		existingTail.next = &newNode
	}

	// Update the list itself
	l.tail = &newNode

	if l.head == nil { // If this is the list's first node
		l.head = &newNode
	}
}

func (l *List) Shift() (interface{}, error) {
	// Update the list
	existingHead := l.head

	switch {
	case l.head.next != nil:
		l.head = existingHead.next
		l.head.prev = nil
	case l.head.next == nil:
		l.head = nil
		l.tail = nil
	}

	// Sever the old head
	existingHead.next = nil
	existingHead.prev = nil

	return existingHead.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	// Update the list
	existingTail := l.tail

	switch {
	case l.tail.prev != nil:
		l.tail = existingTail.prev
		l.tail.next = nil
	case l.tail.prev == nil:
		l.head = nil
		l.tail = nil
	}

	// Sever the old head
	existingTail.next = nil
	existingTail.prev = nil

	return existingTail.Value, nil
}

func (l *List) Reverse() {
	// Initializing iteration
	var currentNode *Node = l.head
	var tempNode *Node = nil

	for currentNode != nil {
		// Intermediate reference so we don't drop a node
		tempNode = currentNode.prev

		// Swap the previous and next
		currentNode.prev = currentNode.next
		currentNode.next = tempNode

		// Update the global to enable iteration
		tempNode = currentNode
		currentNode = currentNode.prev

		/*
		   when we reach the tail, n.next will be nil
		   therefore the swapping of links will yield a prev = nil
		*/
	}

	// swapping original head and tail
	if tempNode != nil {
		l.tail = l.head
		l.head = tempNode
	}
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func NewList(elements ...interface{}) *List {
	linkedList := List{head: nil, tail: nil}

	if len(elements) > 0 {
		linkedList.Unshift(elements[0])
		linkedList.tail = linkedList.head

		for _, value := range elements[1:] {
			linkedList.Push(value)
		}
	}

	return &linkedList
}
