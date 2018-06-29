// Package matrix contains tools to find saddle points in matrices
package matrix

// Pair is a pair of indices in a matrix
type Pair struct {
	i, j int
}

// Saddle returns all Pairs of indices at which the matrix has saddle points
func (m Matrix) Saddle() []Pair {
	var res []Pair

	for i, row := range m.Rows() {
		for j := range row {
			if m.isSaddle(i, j) {
				res = append(res, Pair{i, j})
			}
		}
	}

	return res
}

func (m Matrix) isSaddle(i, j int) bool {
	elem := m[i][j]

	for _, rowEl := range m.Rows()[i] {
		if rowEl > elem {
			return false
		}
	}

	for _, colEl := range m.Cols()[j] {
		if colEl < elem {
			return false
		}
	}

	return true
}
