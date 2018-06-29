// Package meetup implements a simple meetup scheduler
package meetup

import "time"

// WeekSchedule is the type for the description of day we're looking for
type WeekSchedule int

const (
	// First weekday of the month
	First WeekSchedule = iota
	// Second weekday of the month
	Second = iota
	// Third weekday of the month
	Third = iota
	// Fourth weekday of the month
	Fourth = iota
	// Last weekday of the month
	Last = iota
	// Teenth weekday of the month
	Teenth = iota

	searchAscending  = 1
	searchDescending = -1
)

// Day returns the day of the month as described
func Day(schedule WeekSchedule, dow time.Weekday, month time.Month, year int) int {
	firstOfYear := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	switch schedule {
	case First:
		return matchingWeekday(firstOfYear, dow, searchAscending)
	case Second:
		return matchingWeekday(firstOfYear.AddDate(0, 0, 7), dow, searchAscending)
	case Third:
		return matchingWeekday(firstOfYear.AddDate(0, 0, 14), dow, searchAscending)
	case Fourth:
		return matchingWeekday(firstOfYear.AddDate(0, 0, 21), dow, searchAscending)
	case Last:
		return matchingWeekday(firstOfYear.AddDate(0, 1, -1), dow, searchDescending)
	case Teenth:
		return matchingWeekday(firstOfYear.AddDate(0, 0, 12), dow, searchAscending)
	}

	return 0
}

func matchingWeekday(t time.Time, dow time.Weekday, searchDirection int) int {
	for i := 0; i < 7; i++ {
		if t.Weekday() == dow {
			break
		}
		t = t.AddDate(0, 0, searchDirection)
	}

	return t.Day()
}
