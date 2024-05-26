package utils

import (
	"fmt"
	"slices"
	"strings"
)

var cardValues map[byte]int = map[byte]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

var suits []byte = []byte{'C', 'S', 'D', 'H'}

type Card struct {
	value int
	suit  int
}

type Hand struct {
	cards []*Card
}

func NewCard(input string) *Card {
	return &Card{value: cardValues[input[0]], suit: slices.Index(suits, input[1])}
}

func NewHand(input []string) *Hand {
	out := make([]*Card, 0)

	for _, c := range input {
		out = append(out, NewCard(c))
	}

	return &Hand{cards: out}
}

func (hand *Hand) Value() (rank int, value int, kickers []int) {
	// highCard val 	= 1
	// pair val 		= 2
	// 2 pair val		= 3
	// trips val		= 4
	// straight			= 5
	// flush			= 6
	// full house		= 7
	// quads			= 8
	// straight flush	= 9

	cards := hand.cardMap()

	if hand.IsStraight() && hand.IsFlush() {
		max := 0
		for card := range cards {
			if card > max {
				max = card
			}
		}

		return 9, max, nil
	} else if hand.IsQuads() {
		max := 0
		kicker := 0
		for card, v := range cards {
			if v == 4 {
				max = card
			} else {
				kicker = card
			}
		}

		return 8, max, []int{kicker}
	} else if hand.IsFullHouse() {
		trips := 0
		pair := 0
		for card, v := range cards {
			if v == 3 {
				trips = card
			} else {
				pair = card
			}
		}

		return 7, trips, []int{pair}
	} else if hand.IsFlush() {
		max := 0
		for card := range cards {
			if card > max {
				max = card
			}
		}

		return 6, max, nil
	} else if hand.IsStraight() {
		max := 0
		for card := range cards {
			if card > max {
				max = card
			}
		}

		return 5, max, nil
	} else if hand.IsTrips() {
		trips := 0
		kicker1 := 0
		kicker2 := 0
		for card, v := range cards {
			if v == 3 {
				trips = card
			} else if kicker1 == 0 {
				kicker1 = card
			} else {
				kicker2 = card
			}
		}

		if kicker2 > kicker1 {
			kicker1, kicker2 = kicker2, kicker1
		}

		return 4, trips, []int{kicker1, kicker2}
	} else if hand.IsTwoPair() {
		pair1 := 0
		pair2 := 0
		kicker := 0

		for card, v := range cards {
			if v == 2 {
				if pair1 == 0 {
					pair1 = card
				} else {
					pair2 = card
				}
			} else {
				kicker = card
			}
		}

		return 3, Max(pair1, pair2), []int{Min(pair1, pair2), kicker}
	} else if hand.IsOnePair() {
		pair := 0
		kicker1 := 0
		kicker2 := 0
		kicker3 := 0

		for card, v := range cards {
			if v == 2 {
				pair = card
			} else {
				if kicker1 == 0 {
					kicker1 = card
				} else if kicker2 == 0 {
					kicker2 = card
				} else {
					kicker3 = card
				}
			}
		}

		kickers := []int{kicker1, kicker2, kicker3}
		slices.SortFunc(kickers, func(a int, b int) int {
			return b - a
		})

		return 2, pair, kickers
	}

	cardValues := Map(hand.cards, func(card *Card, _ int) int {
		return card.value
	})
	slices.SortFunc(cardValues, func(a int, b int) int {
		return b - a
	})

	return 1, cardValues[0], cardValues[1:]
}

// creates a map of the cards in the hand
//
// e.g. 8C 8S KC 9H 4S =>
// {4: 1, 8: 2, 9: 1, 13: 1}
func (hand *Hand) cardMap() map[int]int {
	out := make(map[int]int)

	for _, card := range hand.cards {
		out[card.value]++
	}

	return out
}

func (hand *Hand) IsStraight() bool {
	cardValues := make([]int, 0)

	for _, card := range hand.cards {
		cardValues = AddIfAbsent(cardValues, card.value)
	}

	if len(cardValues) != 5 {
		return false
	}

	slices.Sort(cardValues)

	for i := cardValues[0]; i <= cardValues[4]; i++ {
		if !slices.Contains(cardValues, i) {
			return false
		}
	}

	return true
}

func (hand *Hand) IsFlush() bool {
	seen := make([]int, 0)
	for _, card := range hand.cards {
		seen = AddIfAbsent(seen, card.suit)
	}

	return len(seen) == 1
}

func (hand *Hand) IsQuads() bool {
	cards := hand.cardMap()

	for _, v := range cards {
		if v == 4 {
			return true
		}
	}

	return false
}

func (hand *Hand) IsFullHouse() bool {
	cards := hand.cardMap()

	if len(cards) != 2 {
		return false
	}

	hasTrips := false
	hasPair := false

	for _, v := range cards {
		if v == 3 {
			hasTrips = true
		}
		if v == 2 {
			hasPair = true
		}
	}

	return hasTrips && hasPair
}

func (hand *Hand) IsTrips() bool {
	cards := hand.cardMap()

	for _, v := range cards {
		if v == 3 {
			return true
		}
	}

	return false
}

func (hand *Hand) IsTwoPair() bool {
	cards := hand.cardMap()

	onePair := false
	for _, v := range cards {
		if v == 2 {
			if !onePair {
				onePair = true
			} else {
				return true
			}
		}
	}

	return false
}

func (hand *Hand) IsOnePair() bool {
	cards := hand.cardMap()

	for _, v := range cards {
		if v == 2 {
			return true
		}
	}

	return false
}

func (hand *Hand) String() string {
	var buf strings.Builder

	for _, card := range hand.cards {
		buf.WriteString(fmt.Sprintf("{suit: %s, value: %d}", string(suits[card.suit]), card.value))
	}

	return buf.String()
}
