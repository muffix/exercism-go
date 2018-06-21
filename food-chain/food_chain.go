// Package foodchain contains tools to construct the lyrics to 'I Know an Old Lady Who Swallowed a Fly'
package foodchain

import (
	"fmt"
	"strings"
)

type verse struct {
	what, comment string
}

var lyrics = []verse{
	{
		what: "fly",
	},
	{
		what:    "spider",
		comment: "It wriggled and jiggled and tickled inside her.",
	},
	{
		what:    "bird",
		comment: "How absurd to swallow a bird!",
	},
	{
		what:    "cat",
		comment: "Imagine that, to swallow a cat!",
	},
	{
		what:    "dog",
		comment: "What a hog, to swallow a dog!",
	},
	{
		what:    "goat",
		comment: "Just opened her throat and swallowed a goat!",
	},
	{
		what:    "cow",
		comment: "I don't know how she swallowed a cow!",
	},
}

// Song returns the entire song
func Song() string {
	return Verses(1, 8)
}

// Verse returns the verse for a given number
func Verse(i int) string {
	if i == 8 {
		return "I know an old lady who swallowed a horse.\nShe's dead, of course!"
	}

	var verse strings.Builder

	verse.WriteString(fmt.Sprintf("%s.\n", lyrics[i-1].what))

	if lyrics[i-1].comment != "" {
		verse.WriteString(fmt.Sprintf("%s\n", lyrics[i-1].comment))
	}

	for j := i - 1; j > 0; j-- {
		victim := lyrics[j-1].what

		if victim == "spider" {
			victim += " that wriggled and jiggled and tickled inside her"
		}

		verse.WriteString(fmt.Sprintf("She swallowed the %s to catch the %s.\n", lyrics[j].what, victim))
	}

	return fmt.Sprintf("I know an old lady who swallowed a %sI don't know why she swallowed the fly. Perhaps she'll die.", verse.String())
}

// Verses returns the verses between the given numbers
func Verses(start, end int) string {
	var verses []string

	for i := start; i <= end; i++ {
		verses = append(verses, Verse(i))

	}

	return strings.Join(verses, "\n\n")
}
