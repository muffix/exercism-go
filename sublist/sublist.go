// Package sublist contains functions to determine the relation between two lists
package sublist

import "reflect"

// Relation is the type of relation of two lists
type Relation string

// Sublist returns the relatuon of the first list to the second (equal, sublist, superlist, unequal)
func Sublist(a, b []int) Relation {
	if len(a) == 0 && len(b) == 0 {
		return "equal"
	}

	var swapped bool

	if len(b) > len(a) {
		a, b = b, a
		swapped = true
	}

	if len(b) == 0 {
		if swapped {
			return "sublist"
		}
		return "superlist"
	}

	for i := 0; i <= len(a)-len(b); i++ {
		if reflect.DeepEqual(a[i:i+len(b)], b) {
			if len(a) == len(b) {
				return "equal"
			}
			if swapped {
				return "sublist"
			}
			return "superlist"
		}
	}

	return "unequal"
}
