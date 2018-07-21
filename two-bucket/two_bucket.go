// Package twobucket implements the two bucket puzzle
package twobucket

import "fmt"

type levels struct {
	one, two int
}
type step struct {
	moves int
	level levels
}

// Solve returns the goal bucket, the number of moves and the level of the other bucket
func Solve(sizeBucketOne,
	sizeBucketTwo,
	goalAmount int,
	startBucket string) (goalBucket string, numSteps, otherBucketLevel int, e error) {

	if sizeBucketOne <= 0 || sizeBucketTwo <= 0 {
		e = fmt.Errorf("Invalid bucket size")
		return
	}

	if goalAmount <= 0 {
		e = fmt.Errorf("Invalid target: %d", goalAmount)
		return
	}

	seen := map[levels]bool{}

	one, two := 0, 0

	if startBucket == "one" {
		seen[levels{0, sizeBucketTwo}] = true
		one, two = sizeBucketOne, 0
	} else if startBucket == "two" {
		seen[levels{sizeBucketOne, 0}] = true
		one, two = 0, sizeBucketTwo
	} else {
		e = fmt.Errorf("Invalid start bucket: %s", startBucket)
		return
	}

	queue := []step{
		{1, levels{one, two}},
	}

	var s step

	for len(queue) > 0 {
		s, queue = queue[0], queue[1:]
		if _, ok := seen[s.level]; ok {
			continue
		}

		seen[s.level] = true

		if s.level.one == goalAmount {
			return "one", s.moves, s.level.two, nil
		}
		if s.level.two == goalAmount {
			return "two", s.moves, s.level.one, nil
		}

		queue = append(queue, step{s.moves + 1, levels{s.level.one, sizeBucketTwo}})
		queue = append(queue, step{s.moves + 1, levels{sizeBucketOne, s.level.two}})
		queue = append(queue, step{s.moves + 1, levels{s.level.one, 0}})
		queue = append(queue, step{s.moves + 1, levels{0, s.level.two}})

		if s.level.one+s.level.two < sizeBucketOne {
			queue = append(queue, step{s.moves + 1, levels{s.level.one + s.level.two, 0}})
		} else {
			queue = append(queue, step{s.moves + 1, levels{sizeBucketOne, s.level.one + s.level.two - sizeBucketOne}})
		}

		if s.level.one+s.level.two < sizeBucketTwo {
			queue = append(queue, step{s.moves + 1, levels{0, s.level.one + s.level.two}})
		} else {
			queue = append(queue, step{s.moves + 1, levels{s.level.one + s.level.two - sizeBucketTwo, sizeBucketTwo}})
		}
	}

	e = fmt.Errorf("No solution")
	return
}
