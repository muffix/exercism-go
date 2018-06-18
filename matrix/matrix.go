// Package matrix contains matrix tools
package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Matrix is the type for a matrix
type Matrix [][]int

// New returns a new matrix parsed from a string
func New(textMatrix string) (Matrix, error) {
	stringRows := strings.Split(textMatrix, "\n")

	var m Matrix

	var dimCols int

	for i, stringRow := range stringRows {
		columns := strings.Split(strings.TrimSpace(stringRow), " ")

		if m == nil {
			dimCols = len(columns)
			m = initMatrix(len(stringRows), dimCols)
		}

		if len(columns) != dimCols {
			return nil, fmt.Errorf("All rows must have the same number of elements")
		}

		for j, value := range columns {
			intValue, err := strconv.Atoi(value)

			if err != nil {
				return nil, err
			}

			m.Set(i, j, intValue)
		}
	}

	return m, nil
}

// Cols returns the columns of a matrix
func (m Matrix) Cols() [][]int {
	dimRows, dimCols := len(m[0]), len(m)
	cols := initMatrix(dimRows, dimCols)

	for i := 0; i < dimRows; i++ {
		for j := 0; j < dimCols; j++ {
			cols.Set(i, j, m[j][i])
		}
	}

	return cols
}

// Rows returns the rows of a matrix
func (m Matrix) Rows() [][]int {
	rows := initMatrix(len(m), len(m[0]))
	for i, row := range m {
		copy(rows[i], row)
	}
	return rows
}

// Set sets the specified matrix entry to the given value
func (m Matrix) Set(row, column, value int) bool {
	if row < 0 || row >= len(m) || column < 0 || column >= len(m[row]) {
		return false
	}

	m[row][column] = value
	return true
}

func initMatrix(rows, columns int) Matrix {
	m := make(Matrix, rows)

	for i := 0; i < rows; i++ {
		m[i] = make([]int, columns)
	}

	return m
}
