// Package strain implements Keep and Discard on collections
package strain

// Ints is the type for a slice of ints
type Ints []int

// Lists is the type for a slice of a slice of ints
type Lists [][]int

// Strings is the type for a slice of ints
type Strings []string

// Keep returns a new Ints with elements from ints for which f evaluates to true
func (ints Ints) Keep(f func(int) bool) Ints {
	if len(ints) == 0 {
		return nil
	}

	filtered := Ints{}

	for _, element := range ints {
		if f(element) {
			filtered = append(filtered, element)
		}
	}

	return filtered
}

// Discard returns a new Ints with elements from ints for which f evaluates to false
func (ints Ints) Discard(f func(int) bool) Ints {
	if len(ints) == 0 {
		return nil
	}

	filtered := Ints{}

	for _, element := range ints {
		if !f(element) {
			filtered = append(filtered, element)
		}
	}

	return filtered
}

// Keep returns a new Lists with elements from lists for which f evaluates to true
func (lists Lists) Keep(f func([]int) bool) Lists {
	if len(lists) == 0 {
		return nil
	}

	filtered := Lists{}

	for _, element := range lists {
		if f(element) {
			filtered = append(filtered, element)
		}
	}

	return filtered
}

// Keep returns a new Strings with elements from strs for which f evaluates to true
func (strs Strings) Keep(f func(string) bool) Strings {
	if len(strs) == 0 {
		return nil
	}

	filtered := Strings{}

	for _, element := range strs {
		if f(element) {
			filtered = append(filtered, element)
		}
	}

	return filtered
}
