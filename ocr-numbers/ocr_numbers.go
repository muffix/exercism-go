// Package ocr implements 7-segment digit recognition
package ocr

import (
	"fmt"
	"reflect"
	"strings"
)

var segmentsInDigit = [10][]int{
	[]int{0, 1, 3, 4, 5, 6},
	[]int{3, 6},
	[]int{0, 2, 3, 4, 5},
	[]int{0, 2, 3, 5, 6},
	[]int{1, 2, 3, 6},
	[]int{0, 1, 2, 5, 6},
	[]int{0, 1, 2, 4, 5, 6},
	[]int{0, 3, 6},
	[]int{0, 1, 2, 3, 4, 5, 6},
	[]int{0, 1, 2, 3, 5, 6},
}

var segmentIndexesInString = [7]int{1, 4, 5, 6, 8, 9, 10}

func recognizeDigit(in string) string {
	var segments []int

	for i, index := range segmentIndexesInString {
		if in[index] != ' ' {
			segments = append(segments, i)
		}
	}

	for digit, digitSegments := range segmentsInDigit {
		if reflect.DeepEqual(digitSegments, segments) {
			return fmt.Sprintf("%d", digit)
		}
	}

	return "?"
}

func recognizeLine(in string) []string {
	lines := strings.Split(in, "\n")

	var out strings.Builder

	for i := 0; i < len(lines[1]); i += 3 {
		var digit strings.Builder

		for line := 0; line <= 2; line++ {
			digit.WriteString(fmt.Sprintln(lines[line][i : i+3]))
		}
		out.WriteString(recognizeDigit(digit.String()))
	}

	return []string{out.String()}
}

// Recognize returns the detected 7-segment digits
func Recognize(in string) []string {
	var out []string

	inLines := strings.Split(in, "\n")
	for i := 1; i < len(inLines); i += 4 {
		lineDigits := recognizeLine(fmt.Sprintf("%s\n%s\n%s\n%s\n", inLines[i], inLines[i+1], inLines[i+2], inLines[i+3]))
		out = append(out, lineDigits[0])
	}

	return out
}
