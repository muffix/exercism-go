// Package palindrome contains tools around number palindromes
package palindrome

import (
	"fmt"
	"sort"
)

type palindromeProducts []Product

// Product is the struct for a product of two integers and its factorizations
type Product struct {
	Product        int
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
}

// Products returns the smallest and largest products which are palindromes from numbers in the given range
func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		return Product{}, Product{}, fmt.Errorf("fmin > fmax is not allowed")
	}

	products := palindromeProductsBetween(fmin, fmax)

	if len(products) == 0 {
		return Product{}, Product{}, fmt.Errorf("no palindromes found in the range between %d and %d", fmin, fmax)
	}

	pmin, pmax = products[0], products[len(products)-1]

	return
}

// palindromeProductsBetween returns all products which are palindromes from numbers in the given range
func palindromeProductsBetween(min, max int) (palindromes palindromeProducts) {
	products := map[int]Product{}

	for i := min; i <= max; i++ {
		for j := i; j <= max; j++ {
			p := i * j

			if !isPalindrome(p) {
				continue
			}

			product, ok := products[p]
			if !ok {
				product = Product{i * j, [][2]int{{i, j}}}
			} else {
				product.Factorizations = append(product.Factorizations, [2]int{i, j})
			}
			products[p] = product
		}
	}

	palindromes = palindromeProducts{}

	for _, product := range products {
		palindromes = append(palindromes, product)
	}

	sort.Sort(palindromes)

	return
}

// isPalindrome returns whether the given number is a palindrome
func isPalindrome(number int) bool {
	str := fmt.Sprintf("%d", number)

	if number < 0 {
		str = str[1:]
	}

	length := len(str)

	for i := 0; i < length; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}

	return true
}

func (pp palindromeProducts) Len() int {
	return len(pp)
}

func (pp palindromeProducts) Swap(i, j int) {
	pp[i], pp[j] = pp[j], pp[i]
}

func (pp palindromeProducts) Less(i, j int) bool {
	return pp[i].Product < pp[j].Product
}
