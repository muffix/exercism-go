package zebra

type Solution struct {
	DrinksWater, OwnsZebra string
}

type constraint func(*house) bool
type house struct {
	colour      Colour
	nationality Nationality
	cigarettes  Cigarettes
	pet         Pet
	drink       Drink
}

type Colour int
type Cigarettes int
type Nationality int
type Pet int
type Drink int

const (
	Red Colour = iota
	Blue
	Yellow
	Green
	Ivory
)

const (
	OldGold Cigarettes = iota
	Parliament
	Kools
	LuckyStrike
	Chesterfield
)

const (
	Norwegian Nationality = iota
	Ukrainian
	Englishman
	Spaniard
	Japanese
)

const (
	Zebra Pet = iota
	Dog
	Horse
	Fox
	Snails
)

const (
	Coffee Drink = iota
	Tea
	Water
	Milk
	OrangeJuice
)

type street []*house

var nationalities = []string{"Norwegian", "Ukranian", "Englishman", "Spaniard", "Japanese"}

var constraints = []constraint{
	func(h *house) bool { // constraint 2
		return (h.nationality == Englishman) == (h.colour == Red)
	},
	func(h *house) bool { // constraint 3
		return (h.nationality == Spaniard) == (h.pet == Dog)
	},
	func(h *house) bool { // constraint 4
		return (h.drink == Coffee) == (h.colour == Green)
	},
	func(h *house) bool { // constraint 5
		return (h.nationality == Ukrainian) == (h.drink == Tea)
	},
	func(h *house) bool { // constraint 7
		return (h.cigarettes == OldGold) == (h.pet == Snails)
	},
	func(h *house) bool { // constraint 8
		return (h.cigarettes == Kools) == (h.colour == Yellow)
	},
	func(h *house) bool { // constraint 13
		return (h.cigarettes == LuckyStrike) == (h.drink == OrangeJuice)
	},
	func(h *house) bool { // constraint 14
		return (h.nationality == Japanese) == (h.cigarettes == Parliament)
	},
}

func SolvePuzzle() Solution {
	street := findStreet()

	sol := Solution{}

	for _, h := range street {
		if h.drink == Water {
			sol.DrinksWater = nationalities[h.nationality]
		}
		if h.pet == Zebra {
			sol.OwnsZebra = nationalities[h.nationality]
		}
	}

	return sol
}

func findStreet() street {
	var houses []*house

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {
				for l := 0; l < 5; l++ {
					for m := 0; m < 5; m++ {
						h := house{
							colour:      Colour(i),
							nationality: Nationality(j),
							cigarettes:  Cigarettes(k),
							pet:         Pet(l),
							drink:       Drink(m),
						}

						if h.matches(constraints) {
							houses = append(houses, &h)
						}
					}
				}
			}
		}
	}

	for i := 0; i < len(houses); i++ {
		if houses[i].nationality != Norwegian { // condition 10
			continue
		}
		for j := 0; j < len(houses); j++ {
			if i == j || houses[j].Conflicts(houses[i]) {
				continue
			}
			for k := 0; k < len(houses); k++ {
				if k == i || k == j || houses[k].Conflicts(houses[i], houses[j]) {
					continue
				}
				if houses[k].drink == Milk { // condition 9
					continue
				}
				for l := 0; l < len(houses); l++ {
					if l == k || l == j || l == i || houses[l].Conflicts(houses[i], houses[j], houses[k]) {
						continue
					}
					for m := 0; m < len(houses); m++ {
						if m == l || m == k || m == j || m == i || houses[m].Conflicts(houses[i], houses[j], houses[k], houses[l]) {
							continue
						}

						candidate := street{houses[i], houses[j], houses[k], houses[l], houses[m]}

						if candidate.isSolution() {
							return candidate
						}
					}
				}
			}
		}
	}
	return nil
}

func (s street) isSolution() bool {
	colPos := make(map[Colour]int, 5)
	natPos := make(map[Nationality]int, 5)
	cigPos := make(map[Cigarettes]int, 5)
	petPos := make(map[Pet]int, 5)
	dnkPos := make(map[Drink]int, 5)

	for i, h := range s {
		colPos[h.colour] = i
		natPos[h.nationality] = i
		cigPos[h.cigarettes] = i
		petPos[h.pet] = i
		dnkPos[h.drink] = i
	}

	if colPos[Green]-colPos[Ivory] != 1 { // condition 6
		return false
	}

	if cigPos[Chesterfield]-petPos[Fox] == 1 || cigPos[Chesterfield]-petPos[Fox] == -1 { // condition 11
		return false
	}

	if cigPos[Kools]-petPos[Horse] == 1 || cigPos[Kools]-petPos[Horse] == -1 { // condition 11
		return false
	}

	if natPos[Norwegian]-colPos[Blue] == 1 || natPos[Norwegian]-colPos[Blue] == -1 { // condition 15
		return false
	}

	return true
}

func (h *house) matches(constraints []constraint) bool {
	for _, validator := range constraints {

		if !validator(h) {
			return false
		}
	}
	return true
}

func (h *house) Conflicts(others ...*house) bool {
	for _, o := range others {
		if h.colour == o.colour || h.nationality == o.nationality || h.cigarettes == o.cigarettes || h.pet == o.pet || h.drink == o.drink {
			return true
		}
	}
	return false
}
