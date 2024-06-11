package problems

import (
	"math"
	"slices"

	"github.com/camnwalter/project-euler/big"
	"github.com/camnwalter/project-euler/utils"
)

func SixtySix() int {
	maxX := big.New(0)
	maxD := 0

	for D := 2; D <= 1000; D++ {
		if utils.IsSquare(D) {
			continue
		}

		fractions := getRepeatedFraction(D)

		periodLength := len(fractions) - 1
		if periodLength%2 == 1 {
			copy := slices.Clone(fractions)
			copy = slices.Delete(copy, 0, 1)
			fractions = append(fractions, copy...)
		}

		a := big.New(0)
		b := big.New(1)
		for i := len(fractions) - 2; i >= 0; i-- {
			b, a = fracAdd(big.New(fractions[i]), a, b)
		}

		if big.Compare(b, maxX) > 0 {
			maxX = b
			maxD = D
		}
	}

	return maxD
}

func getRepeatedFraction(n int) []int {
	if utils.IsSquare(n) {
		return []int{}
	}

	fractions := make([]int, 0)

	sqrt := math.Sqrt(float64(n))
	floor := int(math.Floor(sqrt))

	a := floor
	b := 1
	c := -a

	seen := make([]triple, 0)

	for !slices.Contains(seen, triple{a, b, c}) {
		seen = append(seen, triple{a, b, c})
		fractions = append(fractions, a)

		denom := n - c*c

		num, denom := utils.ReduceFraction(b, denom)
		num *= -c

		a = int(math.Floor((sqrt + float64(num)) / float64(denom)))
		b = denom
		c = -c - denom*a
	}

	return fractions
}
