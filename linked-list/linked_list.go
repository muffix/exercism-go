// Package linkedlist implements a doubly l inked list
package linkedlist

import "fmt"

// ErrEmptyList is an error raised when trying to pop from an empty list
var ErrEmptyList = fmt.Errorf("List is empty")

// Node is a node in a doubly linked list
type Node struct {
	Val            interface{}
	previous, next *Node
}

// List is the type for a duobly-linked list
type List struct {
	head, last *Node
}

// Prev returns the previous node
func (n *Node) Prev() *Node {
	return n.previous
}

// Next returns the next node
func (n *Node) Next() *Node {
	return n.next
}

// First returns the first node
func (n *Node) First() *Node {
	for n.previous != nil {
		n = n.previous
	}
	return n
}

// Last returns the last node
func (n *Node) Last() *Node {
	for n.next != nil {
		n = n.next
	}
	return n
}

// NewList returns a new doubly linked list with the elements
func NewList(elements ...interface{}) *List {
	l := &List{}

	for _, el := range elements {
		l.PushBack(el)
	}

	return l
}

// First returns a pointer to the first node (head)
func (l *List) First() *Node {
	return l.head
}

// Last returns a pointer to the last node
func (l *List) Last() *Node {
	return l.last
}

// PushBack inserts a value at the back
func (l *List) PushBack(v interface{}) {
	node := &Node{Val: v, previous: l.last}

	if l.last == nil {
		l.head = node
	} else {
		l.last.next = node
	}
	l.last = node

}

// PopBack returns and removes a value from the back
func (l *List) PopBack() (interface{}, error) {
	if l.last == nil {
		return 0, ErrEmptyList
	}

	data := l.last.Val

	if l.last.Prev() == nil {
		l.head = nil
		l.last = nil
	} else {
		l.last.Prev().next = nil
		l.last = l.last.Prev()
	}

	return data, nil
}

// PushFront inserts a value at the front
func (l *List) PushFront(v interface{}) {
	node := &Node{Val: v, next: l.head}
	if l.head == nil {
		l.last = node
	} else {
		l.head.previous = node
	}
	l.head = node
}

// PopFront returns and removes a value from the front
func (l *List) PopFront() (interface{}, error) {
	if l.head == nil {
		return 0, ErrEmptyList
	}

	data := l.head.Val

	if l.head.Next() == nil {
		l.head = nil
		l.last = nil
	} else {
		l.head.Next().previous = nil
		l.head = l.head.Next()
	}

	return data, nil

}

// Reverse reverses the list
func (l *List) Reverse() {
	node := l.head
	newHead, newLast := l.last, l.head

	for node != nil {
		node.previous, node.next = node.Next(), node.Prev()
		node = node.Prev()
	}

	l.head, l.last = newHead, newLast
}
