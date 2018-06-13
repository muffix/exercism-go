// Package robotname has functions to handle robot names
package robotname

import (
	"fmt"
	"math/rand"
)

var namesInUse = map[string]bool{}

// Robot is the struct representing the robot
type Robot struct {
	name string
}

// Name returns the robot's name
func (r *Robot) Name() string {
	if r.name == "" {
		nameCandidate := randomName()

		for used := namesInUse[nameCandidate]; used; used = namesInUse[nameCandidate] {
			nameCandidate = randomName()
		}

		r.name = nameCandidate
		namesInUse[nameCandidate] = true
	}
	return r.name
}

// Reset resets the robot's name
func (r *Robot) Reset() {
	r.name = ""
}

// randomName generates a random robot name
func randomName() string {
	return fmt.Sprintf("%s%s%03d", randomLetter(), randomLetter(), rand.Intn(1000))
}

// randomLetter generates a random uppercase letter
func randomLetter() string {
	return string('A' + rune(rand.Intn(26)))
}
