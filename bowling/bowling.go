// Package bowling implements a simple bowling game
package bowling

import (
	"fmt"
)

// Game is the struct representing a bowling game
type Game struct {
	frames []*frame
}

type frame struct {
	rolls, bonusRolls []int
}

// NewGame returns a new game
func NewGame() *Game {
	return &Game{}
}

// Score returns the final score of a game, error if incomplets
func (g *Game) Score() (int, error) {
	if !g.isOver() {
		return 0, fmt.Errorf("Game hasn't finished yet")
	}

	var score int

	for i := 0; i < 10; i++ {
		f := g.frames[i]
		score += f.score()
	}

	return score, nil
}

// Roll records a roll and returns an error for illegal rolls
func (g *Game) Roll(pins int) error {
	if g.isOver() {
		return fmt.Errorf("Game is already over")
	}

	if pins < 0 || pins > 10 {
		return fmt.Errorf("Invalid roll: %d", pins)
	}

	if g.hasOpenFrame() {
		lastFrame := g.frames[len(g.frames)-1]
		if lastFrame.rolls[0]+pins > 10 {
			return fmt.Errorf("Cannot score more than 10 in a frame, got: %d", pins)
		}
		lastFrame.rolls = append(lastFrame.rolls, pins)
	} else {
		g.frames = append(g.frames, &frame{rolls: []int{pins}})
	}

	g.addBonusRolls(pins)

	return nil
}

// isOver returns whether the game is over
func (g *Game) isOver() bool {
	if len(g.frames) < 10 {
		return false
	}

	lastFrame := g.frames[9]
	if len(lastFrame.rolls) < 2 && lastFrame.rolls[0] < 10 {
		return false
	}

	var extraRolls int

	if lastFrame.isStrike() {
		extraRolls = 2
	} else if lastFrame.isSpare() {
		extraRolls = 1
	}

	return extraRolls == len(lastFrame.bonusRolls)
}

// hasOpenFrame returns whether the last frame isn't finished yet
func (g *Game) hasOpenFrame() bool {
	if len(g.frames) == 0 {
		return false
	}
	lastFrame := g.frames[len(g.frames)-1]
	return len(lastFrame.rolls) < 2 && lastFrame.rolls[0] != 10
}

// addBonusRolls adds bonus points to prevuis frames that are eligiblw
func (g *Game) addBonusRolls(pins int) {
	for i := len(g.frames) - 3; i < len(g.frames)-1; i++ {
		if i >= 0 {
			f := g.frames[i]
			if (f.isStrike() && len(f.bonusRolls) < 2) || (f.isSpare() && len(f.bonusRolls) < 1) {
				f.bonusRolls = append(f.bonusRolls, pins)
			}
		}
	}
}

func (f *frame) scoreWithoutBonus() int {
	var score int

	for _, roll := range f.rolls {
		score += roll
	}

	return score
}

func (f *frame) score() int {
	var bonusScore int

	for _, roll := range f.bonusRolls {
		bonusScore += roll
	}

	return bonusScore + f.scoreWithoutBonus()
}

func (f *frame) isStrike() bool {
	return f.scoreWithoutBonus() == 10 && len(f.rolls) == 1
}

func (f *frame) isSpare() bool {
	return f.scoreWithoutBonus() == 10 && len(f.rolls) == 2
}
