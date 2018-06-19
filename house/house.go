package house

import (
	"fmt"
	"strings"
)

var actors = []struct {
	actor  string
	action string
}{
	{
		actor:  "malt",
		action: "lay in",
	},
	{
		actor:  "rat",
		action: "ate",
	},
	{
		actor:  "cat",
		action: "killed",
	},
	{
		actor:  "dog",
		action: "worried",
	},
	{
		actor:  "cow with the crumpled horn",
		action: "tossed",
	},
	{
		actor:  "maiden all forlorn",
		action: "milked",
	},
	{
		actor:  "man all tattered and torn",
		action: "kissed",
	},
	{
		actor:  "priest all shaven and shorn",
		action: "married",
	},
	{
		actor:  "rooster that crowed in the morn",
		action: "woke",
	},
	{
		actor:  "farmer sowing his corn",
		action: "kept",
	},
	{
		actor:  "horse and the hound and the horn",
		action: "belonged to",
	},
}

// Song returns 'This is the House that Jack Built'
func Song() string {
	var song strings.Builder

	for i := 0; i <= len(actors); i++ {
		song.WriteString(fmt.Sprintf("%s\n\n", Verse(i+1)))
	}
	return strings.TrimSpace(song.String())
}

// Verse returns the specified verse of 'This is the House that Jack Built'
func Verse(i int) string {
	var verse strings.Builder
	verse.WriteString("This is the ")

	for ; i > 1; i-- {
		verse.WriteString(fmt.Sprintf("%s\nthat %s the ", actors[i-2].actor, actors[i-2].action))
	}

	verse.WriteString("house that Jack built.")
	return verse.String()
}
