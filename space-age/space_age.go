// Package space contains tools for time calculations on different planets
package space

// Planet is one of the eight planets in the Solar System
type Planet string

const secondsPerEarthYear = 31557600

var orbitalPeriods = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age takes the age in seconds and returns the age on the given planet
func Age(seconds float64, planet Planet) float64 {
	return seconds / secondsPerEarthYear / orbitalPeriods[planet]
}
