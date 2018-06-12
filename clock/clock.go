// Package clock contains functions for time calculation
package clock

import "fmt"

// MinutesPerDay contains the number of minutes in a day
const MinutesPerDay = 24 * 60

// Clock is a type representing a clock
type Clock struct {
	hour, minute int
}

// NewFromMinutes creates a new Clock with the time given in minutes
func NewFromMinutes(minutes int) Clock {
	minutes = minutes % MinutesPerDay
	if minutes < 0 {
		minutes = MinutesPerDay + minutes
	}

	return Clock{minutes / 60, minutes % 60}
}

// New returns a new clock with the given time
func New(hour, minute int) Clock {
	return NewFromMinutes(hour*60 + minute)
}

// Minutes returns a Clock's time in minutes
func (c *Clock) Minutes() int {
	return c.hour*60 + c.minute
}

func (c *Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds a given number of minutes to a given Clock's time and returns a new Clock with the new time
func (c Clock) Add(minutes int) Clock {
	return NewFromMinutes(c.Minutes() + minutes)
}

// Subtract subtracts a given number of minutes from a given Clock's time and returns a new Clock with the new time
func (c Clock) Subtract(minutes int) Clock {
	return c.Add(-minutes)
}
