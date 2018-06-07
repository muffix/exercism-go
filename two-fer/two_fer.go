// Package twofer generates two-fer strings
package twofer

import "fmt"

// ShareWith returns the 2-fer string for the given name
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
