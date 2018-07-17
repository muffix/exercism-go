// Package listops implements basic list operations
package listops

// IntList is the type for a list of integers
type IntList []int
type predFunc func(int) bool
type unaryFunc func(int) int
type binFunc func(int, int) int

// Foldr returns the result of successively applying fn
// to elements of the list, starting from the right
func (l IntList) Foldr(fn binFunc, initial int) int {
	for i := l.Length() - 1; i >= 0; i-- {
		initial = fn(l[i], initial)
	}

	return initial
}

// Foldl returns the result of successively applying fn
// to elements of the list, starting from the left
func (l IntList) Foldl(fn binFunc, initial int) int {
	for i := 0; i < l.Length(); i++ {
		initial = fn(initial, l[i])
	}

	return initial
}

// Filter returns a new listt with elements from l for which
// fn(element) returns true
func (l IntList) Filter(fn predFunc) IntList {
	newList := IntList{}

	for _, el := range l {
		if fn(el) {
			newList = append(newList, el)
		}
	}
	return newList
}

// Length returns the length of l
func (l IntList) Length() int {
	count := 0

	for range l {
		count++
	}
	return count
}

// Reverse returns a new list with elements of l in reverse order
func (l IntList) Reverse() IntList {
	length := l.Length()
	newList := make(IntList, length)
	for i, el := range l {
		newList[length-i-1] = el
	}
	return newList
}

// Map returns a new list with fn applied to every element of l
func (l IntList) Map(fn unaryFunc) IntList {
	length := l.Length()
	newList := make(IntList, length)
	for i, el := range l {
		newList[i] = fn(el)
	}
	return newList
}

// Concat returns a new flat list with elements from all lists
func (l IntList) Concat(others []IntList) IntList {
	target := make(IntList, l.Length())
	copy(target, l)
	for _, other := range others {
		target = target.Append(other)
	}
	return target
}

// Append returns a new list with the contents of l and other
func (l IntList) Append(other IntList) IntList {
	length, otherLength := l.Length(), other.Length()
	newList := make(IntList, length+otherLength)
	copy(newList, l)
	copy(newList[length:], other)
	return newList
}
