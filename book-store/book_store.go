// Package bookstore implements a book store that offers discounts
package bookstore

import (
	"fmt"
	"sort"
)

type basket struct {
	items map[int]int
	cost  int
}
type basketQueue []basket

var discount = []int{
	0, 0, 5, 10, 20, 25,
}

const basePrice = 800

// Cost returns the combined cost for the items
func Cost(items []int) int {
	groups := map[int]int{}
	for _, item := range items {
		groups[item]++
	}

	var lowestCost int
	var b basket
	seen := map[string]int{}

	queue := basketQueue{{groups, 0}}
	for len(queue) > 0 {
		b, queue = queue[0], queue[1:]
		seen[b.itemString()] = b.cost
		if b.isEmpty() && (lowestCost == 0 || b.cost < lowestCost) {
			lowestCost = b.cost
		}

		uniqueItems := b.uniqueItems()
		for i := len(uniqueItems); i > 0; i-- {
			combis := combinations(uniqueItems, i)
			for _, combi := range combis {
				newItems := map[int]int{}
				for item, count := range b.items {
					newItems[item] = count
				}
				for _, item := range combi {
					newItems[item]--
				}
				newCost := b.cost + i*basePrice*(100-discount[i])/100

				if cost, ok := seen[b.itemString()]; !ok || lowestCost == 0 || cost > newCost {
					queue = append(queue, basket{newItems, newCost})
				}
			}
		}
	}

	return lowestCost
}

func (b basket) uniqueItems() []int {
	var unique []int
	for item, count := range b.items {
		if count > 0 {
			unique = append(unique, item)
		}
	}
	sort.Ints(unique)
	return unique
}

func (b basket) isEmpty() bool {
	for _, count := range b.items {
		if count > 0 {
			return false
		}
	}
	return true
}

func (b basket) itemString() string {
	return fmt.Sprintf("%d%d%d%d%d", b.items[1], b.items[2], b.items[3], b.items[4], b.items[5])
}

func combinations(items []int, k int) [][]int {
	combinations := [][]int{}
	if k > len(items) {
		return combinations
	}

	indices := make([]int, k)
	for i := 0; i < k; i++ {
		indices[i] = i
	}
	combinations = append(combinations, subsetFromIndices(items, indices))

	for {
		var i int
		for i = k - 1; i >= 0 && indices[i] == len(items)-k+i; i-- {
		}
		if i < 0 {
			break
		}

		indices[i]++

		i++
		for ; i < k; i++ {
			indices[i] = indices[i-1] + 1
		}
		combinations = append(combinations, subsetFromIndices(items, indices))
	}
	return combinations
}

func subsetFromIndices(original, indices []int) []int {
	subset := make([]int, len(indices))

	for i, index := range indices {
		subset[i] = original[index]
	}
	return subset
}
