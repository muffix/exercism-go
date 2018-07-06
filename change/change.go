// Package change implements a simple change calculator
package change

import "fmt"

type intermediate struct {
	availableCoins, usedCoins []int
	target                    int
}
type queue []intermediate

// Change returns the coins needed to make the target, minimising its number
func Change(coins []int, target int) ([]int, error) {
	q := queue{{coins, []int{}, target}}
	var smallestChange []int

	var i intermediate

	for 0 < len(q) {
		i, q = q[0], q[1:]
		isBetterResult := (len(smallestChange) == 0 || len(i.usedCoins) < len(smallestChange))

		if i.target == 0 && isBetterResult {
			smallestChange = i.usedCoins
		} else if 0 < i.target && 0 < len(i.availableCoins) && isBetterResult {
			coin := i.availableCoins[len(i.availableCoins)-1]
			q = append(q, intermediate{
				availableCoins: i.availableCoins,
				usedCoins:      append([]int{coin}, i.usedCoins...),
				target:         i.target - coin,
			})
			q = append(q, intermediate{
				availableCoins: i.availableCoins[:len(i.availableCoins)-1],
				usedCoins:      i.usedCoins,
				target:         i.target,
			})
		}
	}
	if len(smallestChange) == 0 && target != 0 {
		return nil, fmt.Errorf("Cannot make change %d with coins %v", target, coins)
	}
	return smallestChange, nil
}
