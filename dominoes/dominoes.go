// Package dominoes implements a domino chain finder
package dominoes

// Domino is the type for a domino piece
type Domino [2]int
type queueItem struct {
	chain, remaining []Domino
}

// MakeChain returns whether we can make a chain from the dominoes
func MakeChain(dominoes []Domino) ([]Domino, bool) {
	if len(dominoes) == 0 {
		return dominoes, true
	}

	queue := []queueItem{{[]Domino{dominoes[0]}, dominoes[1:]}}
	var item queueItem

	for len(queue) > 0 {
		item, queue = queue[0], queue[1:]

		lastDomino := item.chain[len(item.chain)-1]
		if len(item.remaining) == 0 && item.chain[0][0] == lastDomino[1] {
			return item.chain, true
		}

		for i, d := range item.remaining {
			if d[0] != lastDomino[1] {
				d[0], d[1] = d[1], d[0]
			}

			if d[0] == lastDomino[1] {
				newChain := make([]Domino, len(item.chain)+1)
				for j := 0; j < len(item.chain); j++ {
					newChain[j] = Domino{item.chain[j][0], item.chain[j][1]}
				}
				newChain[len(item.chain)] = Domino{d[0], d[1]}

				var newRemainder []Domino
				for j, r := range item.remaining {
					if j != i {
						newRemainder = append(newRemainder, Domino{r[0], r[1]})
					}
				}

				queue = append(queue, queueItem{newChain, newRemainder})
			}
		}
	}

	return dominoes, false
}
