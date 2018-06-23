// Package railfence contains tools to decode/encode strings with a rail fence cipher
package railfence

import (
	"strings"
)

type fence [][]rune

type index struct {
	row, column int
}

// readInRowOrder reads a fence in row-first order
func (f fence) readInRowOrder() string {
	var buffer strings.Builder
	for _, index := range f.rowOrderIndexes() {
		buffer.WriteRune(f[index.row][index.column])
	}
	return buffer.String()
}

// readInColumnOrder reads a fence in column-first order
func (f fence) readInColumnOrder() string {
	var buffer strings.Builder

	f.iterate(func(r rune, _, _ int) {
		buffer.WriteRune(r)
	})

	return buffer.String()
}

// iterate iterates a fence in column-first order and passes the current
// rune, row and column to fn.
func (f fence) iterate(fn func(rune, int, int)) {
	direction := 1
	row := 0
	rails := len(f)

	for column := range f[0] {
		fn(f[row][column], row, column)
		if row+direction >= rails || row+direction < 0 {
			direction *= -1
		}
		row += direction
	}
}

// rowOrderIndexes returns an array of indices in row-first order
func (f fence) rowOrderIndexes() []index {
	rows := len(f)

	indexes := make([][]int, rows)
	for i := 0; i < rows; i++ {
		var cols []int
		indexes[i] = cols
	}

	f.iterate(func(_ rune, row, column int) {
		indexes[row] = append(indexes[row], column)
	})

	var flatIndexes []index
	for row, rowIndexes := range indexes {
		for _, column := range rowIndexes {
			flatIndexes = append(flatIndexes, index{row, column})
		}
	}

	return flatIndexes
}

// Encode encodes a message using a rail fence cipher with n rails
func Encode(plain string, n int) string {
	f := makeFence(n, len(plain))
	plainRunes := []rune(plain)

	f.iterate(func(_ rune, row, column int) {
		f[row][column] = plainRunes[column]
	})

	return f.readInRowOrder()
}

// Decode decodes a message using a rail fence cipher with n rails
func Decode(ciphertext string, n int) string {
	fence := makeFence(n, len(ciphertext))
	indexes := fence.rowOrderIndexes()

	for i, r := range ciphertext {
		index := indexes[i]
		fence[index.row][index.column] = r
	}

	return fence.readInColumnOrder()
}

func makeFence(rows, columns int) fence {
	fence := make(fence, rows)

	for i := 0; i < len(fence); i++ {
		fence[i] = make([]rune, columns)
	}

	return fence
}
