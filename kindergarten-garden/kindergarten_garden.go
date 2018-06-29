// Package kindergarten implements a kindergarten garden
package kindergarten

import (
	"fmt"
	"sort"
	"strings"
)

// Garden is the type for the garden
type Garden map[string][]string

var plantSymbols = map[byte]string{
	'C': "clover",
	'G': "grass",
	'R': "radishes",
	'V': "violets",
}

// NewGarden returns a new Garden from the diagram
func NewGarden(diagram string, children []string) (*Garden, error) {
	rows, err := rowsFromDiagram(diagram, len(children))

	if err != nil {
		return nil, err
	}

	g := Garden{}

	for i, child := range sortedCopy(children) {
		if _, ok := g[child]; ok {
			return nil, fmt.Errorf("Cannot have a child name more than once, got: %s", child)
		}

		plants, err := plantsFromRows(rows, i)

		if err != nil {
			return nil, err
		}

		g[child] = plants
	}

	return &g, nil
}

// Plants returns a slice of plants that belong to the child
func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := map[string][]string(*g)[child]
	return plants, ok
}

func sortedCopy(s []string) []string {
	c := make([]string, len(s))
	copy(c, s)
	sort.Strings(c)
	return c
}

func plantsFromRows(rows []string, childIndex int) ([]string, error) {
	plants := make([]string, 4)

	for i := 0; i < 4; i++ {
		symbol := rows[i/2][(childIndex*2)+(i%2)]
		plant, ok := plantSymbols[symbol]

		if !ok {
			return nil, fmt.Errorf("Invalid plant symbol: %s", string(symbol))
		}

		plants[i] = plant
	}

	return plants, nil
}

func rowsFromDiagram(diagram string, childrenCount int) ([]string, error) {
	rows := strings.Split(diagram, "\n")

	if len(rows) != 3 {
		return nil, fmt.Errorf("Must have exactly 3 rows, got %d", len(rows))
	}

	if len(rows[1]) != len(rows[2]) || len(rows[1]) != 2*childrenCount {
		return nil, fmt.Errorf("Diagram doesn't have the required length")
	}

	return rows[1:], nil
}
