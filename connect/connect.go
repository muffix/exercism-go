// Package connect implements the Hex board game
package connect

type board []string

type coordinates struct {
	row, column int
}

type player struct {
	symbol            byte
	startPositions    func(board) []coordinates
	isWinningPosition func(board, coordinates) bool
}

// ResultOf returns the winner of the game
func ResultOf(b board) (string, error) {
	playerX := player{'X', (board).leftEdgeCoordinates, (board).isRightEdge}
	playerO := player{'O', (board).firstRowCoordinates, (board).isLastRow}

	for _, player := range []player{playerO, playerX} {
		visitedCoordinates := map[coordinates]bool{}
		var current coordinates

		for queue := player.startPositions(b); len(queue) > 0; {
			current, queue = queue[0], queue[1:]

			visited := visitedCoordinates[current]
			if visited || b[current.row][current.column] != player.symbol {
				continue
			}

			if player.isWinningPosition(b, current) {
				return string(player.symbol), nil
			}
			visitedCoordinates[current] = true
			queue = append(queue, b.neighbours(current)...)
		}
	}

	return "", nil
}

func (b board) leftEdgeCoordinates() []coordinates {
	var coords []coordinates

	for i := range b[0] {
		coords = append(coords, coordinates{row: i, column: 0})
	}

	return coords
}

func (b board) firstRowCoordinates() []coordinates {
	coords := make([]coordinates, len(b))

	for c := range b[0] {
		coords[c] = coordinates{0, c}
	}

	return coords
}

func (b board) isRightEdge(c coordinates) bool {
	return c.column == len(b[c.row])-1
}

func (b board) isLastRow(c coordinates) bool {
	return c.row == len(b)-1
}

func (b board) neighbours(c coordinates) []coordinates {
	var neighbours []coordinates
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != j {
				cand := coordinates{c.row + i, c.column + j}
				if b.isValidPosition(cand) {
					neighbours = append(neighbours, cand)
				}
			}
		}
	}

	return neighbours
}

func (b board) isValidPosition(c coordinates) bool {
	return c.row >= 0 && c.row < len(b) && c.column >= 0 && c.column < len(b[0])
}
