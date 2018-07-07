// Package spiralmatrix implements a spiral matrix generator
package spiralmatrix

// SpiralMatrix returns a spiral matrix
func SpiralMatrix(n int) [][]int {
	m := make([][]int, n)

	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
	}

	left, top, right, bottom := 0, 0, n-1, n-1

	i := 1

	for left < right {
		// work right, along top
		for c := left; c <= right; c++ {
			m[top][c] = i
			// s[top*n+c] = i
			i++
		}
		top++

		// work down right side
		for r := top; r <= bottom; r++ {
			m[r][right] = i
			i++
		}
		right--
		if top == bottom {
			break
		}

		// work left, along bottom
		for c := right; c >= left; c-- {
			m[bottom][c] = i
			i++
		}
		bottom--

		// work up left side
		for r := bottom; r >= top; r-- {
			m[r][left] = i
			i++
		}
		left++
	}

	// center (last) element
	if n > 0 {
		m[top][left] = i
	}

	return m
}
