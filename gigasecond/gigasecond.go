// Package gigasecond provides time tools in the gigasecond sphere
package gigasecond

import "time"

// AddGigasecond returns the time a gigasecond (10^9 seconds) after the receiver
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
