package problems

import (
	"math"
	"slices"

	"github.com/camnwalter/project-euler/utils"
)

type triple struct {
	a int
	b int
	c int
}

func SixtyFour() int {
	count := 0

	for n := 2; n <= 10_000; n++ {
		if getPeriodLength(n)%2 == 1 {
			count++
		}
	}

	return count
}

func getPeriodLength(n int) int {
	if utils.IsSquare(n) {
		return 0
	}

	sqrt := math.Sqrt(float64(n))
	floor := int(math.Floor(sqrt))

	a := floor
	b := 1
	c := -a

	seen := make([]triple, 0)

	for !slices.Contains(seen, triple{a, b, c}) {
		seen = append(seen, triple{a, b, c})

		denom := n - c*c

		num, denom := utils.ReduceFraction(b, denom)
		num *= -c

		a = int(math.Floor((sqrt + float64(num)) / float64(denom)))
		b = denom
		c = -c - denom*a
	}

	return len(seen) - slices.Index(seen, triple{a, b, c})
}
