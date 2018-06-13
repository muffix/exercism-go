// Package tree contains tools to build a tree
package tree

import (
	"fmt"
	"sort"
)

// Record is the struct to define a node to insert into the tree
type Record struct {
	ID, Parent int
}

// Node is the struct representing a node of the trree
type Node struct {
	ID       int
	Children []*Node
}

// Build creates a tree from the given records and returns the root node
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	if !isContinuousSequence(records) {
		return nil, fmt.Errorf("Records must be a continuous sequence")
	}

	nodes := make(map[int]*Node, cap(records))

	for _, record := range records {
		if record.ID < record.Parent {
			return nil, fmt.Errorf("Children can't have parents younger than themselves")
		}

		if _, ok := nodes[record.ID]; !ok {
			nodes[record.ID] = &Node{ID: record.ID}

			if parent, ok := nodes[record.Parent]; ok && record.ID != record.Parent {
				parent.Children = append(parent.Children, nodes[record.ID])
			} else {
				if record.ID != 0 {
					return nil, fmt.Errorf("Parent %d can't be younger than the child %d", record.Parent, record.ID)
				}
			}
		} else {
			return nil, fmt.Errorf("Node with ID %d already exists", record.ID)
		}
	}

	if len(nodes) != len(records) {
		return nil, fmt.Errorf("Generated tree has %d nodes, expected %d", len(nodes), len(records))
	}

	return nodes[0], nil
}

func isContinuousSequence(records []Record) bool {
	for i, record := range records {
		if record.ID != i {
			return false
		}
	}
	return true
}
