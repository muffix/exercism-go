// Package encode implements run-length encoding and decoding
package encode

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// RunLengthEncode returns the encoded string
func RunLengthEncode(in string) string {
	var series byte
	var index, count int
	var encoded strings.Builder

	for index < len(in) {
		series = in[index]
		count = 0
		for index < len(in) && in[index] == series {
			count++
			index++
		}

		if count > 1 {
			encoded.WriteString(fmt.Sprintf("%d", count))
		}
		encoded.WriteByte(series)
	}

	return encoded.String()
}

// RunLengthDecode returns the decoded string
func RunLengthDecode(in string) string {
	var decoded strings.Builder

	matches := regexp.MustCompile(`(\d+)?([^\d])`).FindAllStringSubmatch(in, -1)
	for _, match := range matches {
		count, err := strconv.Atoi(match[1])
		if err != nil {
			count = 1
		}
		decoded.WriteString(strings.Repeat(match[2], count))
	}

	return decoded.String()
}
