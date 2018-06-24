// Package stringset contains a set type that holds strings
package stringset

import (
	"fmt"
	"strings"
)

// Set represents the datatype
type Set map[string]struct{}

var sentinel struct{}

// New returns a new Set
func New() Set {
	return Set{}
}

// NewFromSlice returns a new Set with elements from the slice
func NewFromSlice(slice []string) Set {
	set := Set{}

	for _, str := range slice {
		set.Add(str)
	}

	return set
}

func (s Set) String() string {
	var elements []string

	for element := range s {
		elements = append(elements, fmt.Sprintf("\"%s\"", element))
	}

	return fmt.Sprintf("{%s}", strings.Join(elements, ", "))
}

// IsEmpty returns whether the Set is empty
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// Has returns whether the Set contains the string
func (s Set) Has(str string) bool {
	_, ok := s[str]
	return ok
}

// Subset returns whether s1 is a subset of s2
func Subset(s1, s2 Set) bool {
	for element := range s1 {
		if !s2.Has(element) {
			return false
		}
	}
	return true
}

// Disjoint returns whether the sets have no single element in common
func Disjoint(s1, s2 Set) bool {
	for element := range s1 {
		if s2.Has(element) {
			return false
		}
	}
	return true
}

// Equal returns whether the sets are equal
func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	return Subset(s1, s2)
}

// Add adds an element to a set
func (s Set) Add(str string) {
	s[str] = sentinel
}

// Intersection returns the intersection of the two sets
func Intersection(s1, s2 Set) Set {
	s := New()

	for element := range s1 {
		if s2.Has(element) {
			s.Add(element)
		}
	}

	for element := range s2 {
		if s1.Has(element) {
			s.Add(element)
		}
	}

	return s
}

// Difference returns a new Set with the elements from s1 which are not in s2
func Difference(s1, s2 Set) Set {
	s := New()

	for element := range s1 {
		if !s2.Has(element) {
			s.Add(element)
		}
	}
	return s
}

// Union returns the union of the sets
func Union(s1, s2 Set) Set {
	s := New()
	for elem := range s1 {
		s.Add(elem)
	}
	for elem := range s2 {
		s.Add(elem)
	}
	return s
}
