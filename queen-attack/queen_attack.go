//Package queenattack contains tools to determine whether two queens on a chess board can attack each other
package queenattack

import (
	"fmt"
)

// CanQueenAttack returns whether two queens on a chess board can attack each other
func CanQueenAttack(w, b string) (bool, error) {
	if w == b {
		return false, fmt.Errorf("Two pieces can't be in the same position: %s", w)
	}

	whiteX, whiteY, err := intPosition(w)

	if err != nil {
		return false, err
	}

	blackX, blackY, err := intPosition(b)

	if err != nil {
		return false, err
	}

	// If they're in either the same row or the same column, they can attack
	if whiteX == blackX || whiteY == blackY {
		return true, nil
	}

	// Check the diagonals in all 4 directions as long as we're not leaving the board
	for _, dX := range []int{-1, 1} {
		for _, dY := range []int{-1, 1} {
			x, y := whiteX, whiteY
			for i := 0; withinBounds(x, y); i++ {
				// if board[x][y] && x != whiteX {
				if x == blackX && y == blackY && x != whiteX {
					return true, nil
				}
				x += dX
				y += dY
			}
		}
	}

	return false, nil
}

// intPosition returns 0-indexed integers for the position of the piece
func intPosition(pos string) (int, int, error) {
	if len(pos) != 2 {
		return -1, -1, fmt.Errorf("Invalid position")
	}

	x := int(pos[0] - 'a')
	y := int(pos[1] - 49)

	if !withinBounds(x, y) {
		return -1, -1, fmt.Errorf("Invalid position")
	}

	return x, y, nil
}

// withinBounds returns whether a position is within the limits of the chess coard
func withinBounds(x, y int) bool {
	return !(x < 0 || x > 7 || y < 0 || y > 7)
}
