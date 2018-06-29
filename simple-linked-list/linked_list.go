// Package linkedlist implements a simple linked list
package linkedlist

import (
	"fmt"
)

// Element is an element of a linked list
type Element struct {
	data int
	next *Element
}

// List is the type for a linked list
type List struct {
	head *Element
	size int
}

// New returns a new linked list
func New(ints []int) *List {
	l := &List{}

	for _, i := range ints {
		l.Push(i)
	}

	return l
}

// Size returns the size of the lsit
func (l *List) Size() int {
	return l.size
}

// Push adds an element to the end of the list
func (l *List) Push(i int) {
	if l.head == nil {
		l.head = &Element{data: i}
		l.size = 1

		return
	}

	element := l.head

	for element.next != nil {
		element = element.next
	}

	element.next = &Element{data: i}
	l.size++
}

// Pop returns the data of the last element of the list and removes it from the list
func (l *List) Pop() (int, error) {
	if l.head == nil {
		return 0, fmt.Errorf("List is empty")
	}

	if l.head.next == nil {
		res := l.head.data

		l.head = nil
		l.size = 0

		return res, nil
	}

	element := l.head

	for element.next.next != nil {
		element = element.next
	}

	res := element.next.data

	element.next = nil
	l.size--

	return res, nil
}

// Array returns the data of the list in a slice
func (l *List) Array() []int {
	ints := []int{}

	element := l.head

	for element != nil {
		ints = append(ints, element.data)
		element = element.next
	}

	return ints
}

// Reverse returns a new list in reverse order
func (l *List) Reverse() *List {
	element := l.head
	ints := []int{}

	for element != nil {
		ints = append([]int{element.data}, ints...)
		element = element.next
	}

	return New(ints)
}
