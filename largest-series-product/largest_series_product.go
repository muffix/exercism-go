// Package lsproduct contains tools to calculate the largest product for a contiguous substring of digits of length n
package lsproduct

import (
	"fmt"
	"strconv"
)

// LargestSeriesProduct calculate the largest product for a contiguous substring of digits of length n
func LargestSeriesProduct(digits string, n int) (int64, error) {
	if n < 0 || n > len(digits) {
		return -1, fmt.Errorf("Invalid length of series:  %d", n)
	}

	var maxProduct int64

	for i := 0; i <= len(digits)-n; i++ {

		current, err := sumBetween(i, i+n, digits)

		if err != nil {
			return -1, err
		}

		if current > maxProduct {
			maxProduct = current
		}
	}

	return maxProduct, nil
}

func sumBetween(start, end int, digits string) (int64, error) {
	var res int64 = 1

	for i := start; i < end; i++ {

		digit, err := strconv.Atoi(string(digits[i]))

		if err != nil {
			return -1, err
		}

		res *= int64(digit)
	}

	return res, nil
}
