// Package binarysearch implements a binary search on a slice of ints
package binarysearch

// SearchInts returns the index of the key in the slice
func SearchInts(slice []int, key int) int {

	lo, hi := 0, len(slice)-1

	for lo < hi {
		split := (lo + hi) / 2

		if slice[split] < key {
			lo = split + 1
		} else {
			hi = split
		}
	}

	if lo == len(slice) || slice[lo] != key {
		return -1
	}

	return lo
}
