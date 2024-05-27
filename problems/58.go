package problems

import (
	"math"

	"github.com/camnwalter/project-euler/utils"
)

func FiftyEight() int {
	a, b, c, d := 1, 1, 1, 1
	n := 2

	found := 0
	diagonals := 1

	for step := 1; found == 0 || float64(found)/float64(diagonals) >= 0.1; step++ {
		a += n
		n += 2

		b += n
		n += 2

		c += n
		n += 2

		d += n
		n += 2

		if utils.IsPrime(a) {
			found++
		}

		if utils.IsPrime(b) {
			found++
		}

		if utils.IsPrime(c) {
			found++
		}

		diagonals += 4
	}

	return int(math.Sqrt(float64(d)))
}
