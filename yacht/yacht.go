// Package yacht implements the dice game Yacht.
package yacht

type scorer func(roll) int

type roll map[int]int

var categories = map[string]scorer{
	"ones":            fixedNum(1),
	"twos":            fixedNum(2),
	"threes":          fixedNum(3),
	"fours":           fixedNum(4),
	"fives":           fixedNum(5),
	"sixes":           fixedNum(6),
	"full house":      fullHouse,
	"four of a kind":  fourOfAKind,
	"little straight": straight(1),
	"big straight":    straight(2),
	"yacht":           yacht,
	"choice":          sumOfDice,
}

// Score returns the score of the dice roll in the category.
func Score(dice []int, cat string) int {
	r := roll{}

	for _, d := range dice {
		r[d]++
	}

	return categories[cat](r)
}

func fixedNum(num int) func(roll) int {
	return func(r roll) int {
		return r[num] * num
	}
}

func fullHouse(r roll) int {
	var foundTwo, foundThree bool
	for _, count := range r {
		switch count {
		case 2:
			foundTwo = true
		case 3:
			foundThree = true
		}
	}

	if foundTwo && foundThree {
		return sumOfDice(r)
	}

	return 0
}

func fourOfAKind(r roll) int {
	for val, count := range r {
		if count >= 4 {
			return val * 4
		}
	}
	return 0
}

func straight(start int) func(roll) int {
	return func(r roll) int {
		for i := start; i <= start+4; i++ {
			if count := r[i]; count != 1 {
				return 0
			}
		}
		return 30
	}
}

func yacht(r roll) int {
	if len(r) == 1 {
		return 50
	}
	return 0
}

func sumOfDice(r roll) int {
	var score int

	for v, count := range r {
		score += v * count
	}

	return score
}
