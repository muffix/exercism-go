// Package twelve contains tools to print the song 'The Twelve Days of Christmas'
package twelve

import (
	"bytes"
	"fmt"
)

type verse struct {
	day, present string
}

var lyrics = [12]verse{
	{"first", "a Partridge in a Pear Tree"},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}

// Song returns the entire song
func Song() string {
	var verses bytes.Buffer
	for i := 1; i <= 12; i++ {
		verses.WriteString(fmt.Sprintln(Verse(i)))
	}
	return verses.String()
}

// Verse returns the verse for a given day
func Verse(i int) string {
	var verse bytes.Buffer

	for j := i - 1; j > 0; j-- {
		verse.WriteString(fmt.Sprintf("%s, ", lyrics[j].present))
	}

	if i > 1 {
		verse.WriteString("and ")
	}

	verse.WriteString(lyrics[0].present)

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me, %s.", lyrics[i-1].day, verse.String())
}
