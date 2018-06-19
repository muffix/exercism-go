// Package pascal contains functions to compute Pascal triangles
package pascal

// Triangle computes the Pascal triangle up to the given level
func Triangle(targetLevel int) (triangle [][]int) {
	triangle = make([][]int, targetLevel)

	for i := range triangle {
		for j := 0; j <= i; j++ {
			triangle[i] = append(triangle[i], binom(i, j))
		}
	}

	return triangle
}

func binom(n, k int) int {
	res := 1
	if k > n-k {
		k = n - k
	}
	for i := 0; i < k; i++ {
		res *= n - i
		res /= i + 1
	}
	return res
}
