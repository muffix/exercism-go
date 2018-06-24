// Package say contains tools to generate the natural language representation of a number
package say

import (
	"fmt"
	"strings"
)

var firstNineteen = [20]string{
	"",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}

var tens = [8]string{
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}

var higher = [4]string{
	"",
	"thousand",
	"million",
	"billion",
}

func nhundred(n int) string {
	if n == 0 {
		return ""
	}
	return fmt.Sprintf("%s hundred ", firstNineteen[n])
}

func sayLessThanOneHundred(n int) string {
	if n < 20 {
		return firstNineteen[n]
	}

	if n%10 == 0 {
		return tens[(n/10)-2]
	}

	return fmt.Sprintf("%s-%s", tens[(n/10)-2], firstNineteen[n%10])

}

func sayLessThanOneThousand(n int) string {
	if n == 0 {
		return ""
	}

	var number strings.Builder

	number.WriteString(nhundred(int((n / 100) % 10)))
	number.WriteString(sayLessThanOneHundred(int(n % 100)))

	return number.String()
}

// Say returns the natural language representaion of n
func Say(n int64) (string, bool) {
	if n < 0 || n >= 1000000000000 {
		return "", false
	}
	if n == 0 {
		return "zero", true
	}

	var blocks []string

	for exp := 0; n > 0; exp += 3 {
		current := sayLessThanOneThousand(int(n % 1000))
		if current != "" {
			blocks = append([]string{current, higher[exp/3]}, blocks...)
		}
		n /= 1000
	}

	return strings.TrimSpace(strings.Join(blocks, " ")), true
}
