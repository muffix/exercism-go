// Package anagram contains tools to detect anagrams
package anagram

import (
	"reflect"
	"sort"
	"strings"
)

type sortableString []byte

func (s sortableString) Len() int {
	return len(s)
}

func (s sortableString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortableString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Detect returns all case-insensitive anagrams of the subject among the candidates.
func Detect(subject string, candidates []string) []string {
	subject = strings.ToLower(subject)
	target := sortableString(subject)
	sort.Sort(target)

	var detected []string

	for _, candidate := range candidates {
		if strings.ToLower(candidate) == subject {
			continue
		}

		c := sortableString(strings.ToLower(candidate))
		sort.Sort(c)

		if reflect.DeepEqual(c, target) {
			detected = append(detected, candidate)
		}
	}

	return detected
}
