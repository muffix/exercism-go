// Package pov implements reparenting a graph on a selected node
package pov

import "fmt"

// Graph is the type for a directed graph
type Graph map[string]map[string]bool

// New returns a new Graph
func New() *Graph {
	return &Graph{}
}

// AddNode adds a node with the label to the graph
func (g *Graph) AddNode(nodeLabel string) {
	(*g)[nodeLabel] = map[string]bool{}
}

// AddArc adds a directed arc between from and to to the graph
func (g *Graph) AddArc(from, to string) {
	if (*g)[from] == nil {
		(*g)[from] = map[string]bool{}
	}

	(*g)[from][to] = true
}

// ArcList returns a slice of string representations of arcs in the graph
func (g *Graph) ArcList() []string {
	arcStrings := []string{}
	for from, arcs := range *g {
		for to := range arcs {
			arcStrings = append(arcStrings, fmt.Sprintf("%s -> %s", from, to))
		}
	}
	return arcStrings
}

// ChangeRoot reparents the graph to the new root
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	path := g.findPath(oldRoot, newRoot)

	for i := 0; i < len(path)-1; i++ {
		to, from := path[i], path[i+1]

		delete((*g)[from], to)
		g.AddArc(to, from)
	}

	return g
}

// findPath returns a path between the nodes from and to
func (g *Graph) findPath(from, to string) []string {
	if from == to {
		return []string{to}
	}

	for child := range (*g)[from] {
		if path := g.findPath(child, to); path != nil {
			return append(path, from)
		}
	}

	return nil

}
