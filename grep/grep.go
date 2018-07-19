// Package grep implements grep
package grep

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type match struct {
	lineNo int
	line   string
}

type fileMatch struct {
	fileName string
	matches  []match
}

const (
	flagN = 2 << iota
	flagL = 2 << iota
	flagI = 2 << iota
	flagV = 2 << iota
	flagX = 2 << iota
)

// Search returns the search restults for pattern in the given files
func Search(pattern string, flagsStrings, files []string) []string {
	flags := parseFlags(flagsStrings)

	var allMatches []fileMatch
	for _, fileName := range files {
		var matchesInFile []match
		content, _ := ioutil.ReadFile(fileName)
		lines := strings.Split(string(content), "\n")

		for i, line := range lines {
			if matches(pattern, flags, line) && line != "" {
				matchesInFile = append(matchesInFile, match{i + 1, line})
				if flags&flagL > 0 {
					break
				}
			}
		}

		allMatches = append(allMatches, fileMatch{fileName, matchesInFile})

	}

	return formatMatches(allMatches, flags)
}

func parseFlags(flagStrings []string) int {
	flags := 0

	for _, f := range flagStrings {
		switch f {
		case "-i":
			flags |= flagI
		case "-l":
			flags |= flagL
		case "-n":
			flags |= flagN
		case "-v":
			flags |= flagV
		case "-x":
			flags |= flagX
		}
	}

	return flags
}

func matches(pattern string, flags int, text string) bool {
	if flags&flagI > 0 {
		pattern, text = strings.ToLower(pattern), strings.ToLower(text)
	}

	if flags&flagX > 0 {
		return pattern == strings.TrimSpace(text)
	} else if flags&flagV > 0 {
		return !strings.Contains(text, pattern)
	}

	return strings.Contains(text, pattern)
}

func formatMatches(fileMatches []fileMatch, flags int) []string {
	matchesResult := []string{}

	for _, fm := range fileMatches {
		for _, m := range fm.matches {
			var match string

			if flags&flagL > 0 {
				match = fmt.Sprintf(fm.fileName)
			} else if flags&flagN > 0 {
				match = fmt.Sprintf("%d:%s", m.lineNo, m.line)
			} else {
				match = fmt.Sprintf(m.line)
			}

			if len(fileMatches) > 1 && flags&flagL == 0 {
				match = fmt.Sprintf("%s:%s", fm.fileName, match)
			}

			matchesResult = append(matchesResult, match)

			if flags&flagL > 0 {
				break
			}
		}
	}

	return matchesResult
}
