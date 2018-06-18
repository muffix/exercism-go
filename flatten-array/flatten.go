// Package flatten contains tools to flatten nested slices
package flatten

// Flatten returns a flat slice from the input
func Flatten(in interface{}) []interface{} {
	out := []interface{}{}

	switch element := in.(type) {
	case []interface{}:
		for _, el := range element {
			out = append(out, Flatten(el)...)
		}
	case interface{}:
		out = append(out, element)
	}

	return out
}
