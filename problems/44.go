package problems

import (
	"fmt"
	"math"
)

func FortyFour() {
	var difference int

	for i := 1; i < 10_000; i++ {
		for j := i + 1; j < 10_000; j++ {
			p_i := i * (3*i - 1) / 2
			p_j := j * (3*j - 1) / 2

			sum := p_i + p_j
			diff := p_j - p_i

			if IsPentagonal(sum) && IsPentagonal(diff) {
				difference = diff
				break
			}
		}
	}

	fmt.Println("min pentagonal diff =", difference)
}

func IsPentagonal(p int) bool {
	// p = n (3n - 1) / 2
	// 1.5n^2 - .5n - p = 0
	// n = (1 + sqrt(1 + 24p)) / 6

	n := (1 + math.Sqrt(float64(1+24*p))) / 6
	return math.Round(n) == n
}
