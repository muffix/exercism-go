package accumulate

// Accumulate returns a new slice by applying the given converter to every element in the input
func Accumulate(input []string, converter func(string) string) []string {
	var output = make([]string, len(input))
	for i, str := range input {
		output[i] = converter(str)
	}
	return output
}
