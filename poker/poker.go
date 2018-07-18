// Package poker determines the highest value hand(s) from a slice of strings
package poker

import (
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strconv"
)

type pokerHand int

const (
	HandStraightFlush pokerHand = iota
	HandFourOfAKind
	HandFullHouse
	HandFlush
	HandStraight
	HandThreeOfAKind
	HandTwoPair
	HandOnePair
	HandHighCard
)

type hand struct {
	original string
	ranks    []int
	groups   rankGroups
	flush    string
}

type rankGroup struct {
	count, rank int
}
type rankGroups []rankGroup

type hands []hand

func (h *hand) nTuples(n int) int {
	var tuples int
	for _, count := range h.ranks {
		if count == n {
			tuples++
		}
	}
	return tuples
}

var cardRegex = regexp.MustCompile(`([2-9JQKA]|10)([♧♤♡♢])(?: |$)`)

// BestHand finds the best hand(s) from the ones given in the string slice
func BestHand(handStrings []string) ([]string, error) {
	var playerHands hands

	for _, hand := range handStrings {
		h, err := newHand(hand)
		if err != nil {
			return nil, err
		}
		playerHands = append(playerHands, h)
	}

	sort.Sort(playerHands)

	return playerHands.top(), nil
}

// highestHand returns the highest value of the hand
func (h *hand) highestHand() pokerHand {
	switch {
	case h.flush != "" && h.isStraight():
		return HandStraightFlush
	case h.groups[0].count == 4:
		return HandFourOfAKind
	case h.groups[0].count == 3 && h.groups[1].count == 2:
		return HandFullHouse
	case h.flush != "":
		return HandFlush
	case h.isStraight():
		return HandStraight
	case h.groups[0].count == 3:
		return HandThreeOfAKind
	case h.groups[0].count == 2 && h.groups[1].count == 2:
		return HandTwoPair
	case h.groups[0].count == 2:
		return HandOnePair
	default:
		return HandHighCard
	}
}

// isStraight returns whether the hand is a straight
// If an ace is played low, it sorts the hand accordingly.
func (h *hand) isStraight() bool {
	ranks := h.ranks

	if ranks[0] == 14 && ranks[1] == 5 {
		ranks = append(ranks[1:], 1)
	}

	for i := 1; i < 5; i++ {
		if ranks[i] != ranks[i-1]-1 {
			return false
		}
	}
	h.ranks = ranks
	return true
}

// newHand parses s and returns a new hand with sorted ranks and groups
func newHand(s string) (hand, error) {
	matches := cardRegex.FindAllStringSubmatch(s, -1)

	if len(matches) != 5 {
		return hand{}, fmt.Errorf("Invalid hand: %s", s)
	}

	h := hand{original: s}
	flush := true

	cards := map[int]int{}

	for _, card := range matches {
		var rank int
		switch card[1] {
		case "A":
			rank = 14
		case "K":
			rank = 13
		case "Q":
			rank = 12
		case "J":
			rank = 11
		default:
			rank, _ = strconv.Atoi(card[1])
		}
		cards[rank]++
		flush = flush && card[2] == matches[1][2]
	}

	for rank, count := range cards {
		h.groups = append(h.groups, rankGroup{count, rank})
	}

	sort.Sort(sort.Reverse(h.groups))
	for _, g := range h.groups {
		for i := 0; i < g.count; i++ {
			h.ranks = append(h.ranks, g.rank)
		}
	}

	if flush {
		h.flush = matches[0][2]
	}

	return h, nil
}

// top returns the best hand, possibly more than one if they're equal
func (h hands) top() (topHands []string) {
	topHands = []string{h[0].original}

	if len(h) == 1 {
		return
	}

	for i := 1; i < len(h); i++ {
		if reflect.DeepEqual(h[i].ranks, h[0].ranks) && h[i].highestHand() == h[0].highestHand() {
			topHands = append(topHands, h[i].original)
			continue
		}
		break
	}
	return
}

func (h hands) Len() int { return len(h) }

func (h hands) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h hands) Less(i, j int) bool {
	if h[i].highestHand() == h[j].highestHand() {
		for k := 0; k < 5; k++ {
			if h[i].ranks[k] != h[j].ranks[k] {
				return h[i].ranks[k] > h[j].ranks[k]
			}
		}
	}
	return h[i].highestHand() < h[j].highestHand()
}

func (g rankGroups) Len() int { return len(g) }

func (g rankGroups) Swap(i, j int) { g[i], g[j] = g[j], g[i] }

func (g rankGroups) Less(i, j int) bool {
	if g[i].count == g[j].count {
		return g[i].rank <= g[j].rank
	}
	return g[i].count < g[j].count
}
