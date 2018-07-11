// Package alphametics implements a solver for alphametics puzzles
package alphametics

import (
	"fmt"
	"sort"
	"strings"

	"github.com/cznic/mathutil"
)

// Solve returns an assignment of values to letters that solve the puzzle
func Solve(puzzle string) (map[string]int, error) {
	components := strings.Split(puzzle, " == ")
	left, target := components[0], components[1]

	summands := strings.Split(left, " + ")

	lettersInUse := map[string]bool{}
	var lettersWithOrder []string

	for _, r := range puzzle {
		if r >= 'A' && r <= 'Z' {
			letter := string(r)
			if _, used := lettersInUse[letter]; !used {
				lettersWithOrder = append(lettersWithOrder, letter)
			}
			lettersInUse[letter] = true
		}
	}

	values := make(chan map[string]int)

	go makeValues(lettersWithOrder, values)

	var v map[string]int

	for v = range values {
		if verify(summands, target, v) {
			return v, nil
		}
	}

	return nil, fmt.Errorf("No solution found")
}

// makeValues generates all the value assignments we need to test
func makeValues(letters []string, values chan<- map[string]int) {
	combiCh := make(chan []int)
	permCh := make(chan []int)
	go combinations(10, len(letters), combiCh)
	go permutations(combiCh, permCh)

	for permutation := range permCh {
		vals := make(map[string]int, len(letters))
		for i := 0; i < len(letters); i++ {
			vals[letters[i]] = permutation[i]
		}
		values <- vals
	}

	close(values)
}

// verify returns whether a solution is valid
func verify(summands []string, target string, values map[string]int) bool {
	var sum int

	// invalid if the first number is a 0 for the target
	if firstVal := values[string(target[0])]; firstVal == 0 {
		return false
	}

	for _, summand := range summands {
		sum += value(summand, values)
		// invalid if the first number of any summand is a 0
		if firstVal := values[string(summand[0])]; firstVal == 0 {
			return false
		}
	}

	targetValue := value(target, values)

	// valid if the summands sum up to the target
	return sum == targetValue
}

// value returns the value of a number with alphadigits with the given valued
func value(n string, values map[string]int) int {
	var total int
	exp := 1
	for i := len(n) - 1; i >= 0; i-- {
		total += values[string(n[i])] * exp
		exp *= 10
	}

	return total
}

// permutations sends all permutations of the slices it reads from combiCh to permCh
func permutations(combiCh <-chan []int, permCh chan<- []int) {

	for slice := range combiCh {
		sortableSlice := sort.IntSlice(slice)

		mathutil.PermutationFirst(sortableSlice)
		permCh <- sortableSlice

		for mathutil.PermutationNext(sortableSlice) {
			perm := make([]int, len(sortableSlice))
			copy(perm, sortableSlice)
			permCh <- perm
		}
	}

	close(permCh)
}

// combinations generates all k-combinations of the first n numbers
func combinations(n, k int, ch chan<- []int) {
	if k > n {
		close(ch)
	}

	combination := make([]int, k)
	for i := 0; i < k; i++ {
		combination[i] = i
	}
	newCombi := make([]int, k)
	copy(newCombi, combination)
	ch <- newCombi

	for {
		var i int
		for i = k - 1; i >= 0 && combination[i] == n-k+i; i-- {
		}
		if i < 0 {
			break
		}

		combination[i]++

		i++
		for ; i < k; i++ {
			combination[i] = combination[i-1] + 1
		}
		newCombi := make([]int, k)
		copy(newCombi, combination)
		ch <- newCombi
	}
	close(ch)
}
