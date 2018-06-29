// Package binarysearchtree implements a binary search tree
package binarysearchtree

// SearchTreeData is the struct for a node of the tree
type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// Bst returns a new binary search tree with the given value at the root
func Bst(root int) *SearchTreeData {
	return &SearchTreeData{data: root}
}

// Insert inserts a new node into the tree
func (node *SearchTreeData) Insert(i int) {
	if node.data >= i {
		if node.left != nil {
			node.left.Insert(i)
		} else {
			node.left = &SearchTreeData{data: i}
		}
	} else if node.data < i {
		if node.right != nil {
			node.right.Insert(i)
		} else {
			node.right = &SearchTreeData{data: i}
		}
	}
}

// MapString applies the function to every element and returns the results in order
func (node *SearchTreeData) MapString(f func(int) string) []string {
	if node == nil {
		return []string{}
	}
	result := append(node.left.MapString(f), f(node.data))
	result = append(result, node.right.MapString(f)...)
	return result
}

// MapInt applies the function to every element and returns the results in order
func (node *SearchTreeData) MapInt(f func(int) int) []int {
	if node == nil {
		return []int{}
	}
	result := append(node.left.MapInt(f), f(node.data))
	result = append(result, node.right.MapInt(f)...)
	return result
}
