package problems

import (
	"fmt"
	"math"
)

func ThirtyNine() {
	maxSolutions := 0
	maxP := 0

	for p := 1; p <= 1000; p++ {
		solutions := 0

		for a := 1; a < p; a++ {
			for b := 1; b < p-a; b++ {
				if yes, c := IsPerfectSquare(a*a + b*b); yes && a+b+c == p {
					solutions++
				}
			}
		}

		if solutions > maxSolutions {
			maxSolutions = solutions
			maxP = p
		}
	}

	fmt.Println("max solutions =", maxP)
}

func IsPerfectSquare(n int) (bool, int) {
	sqrt := math.Sqrt(float64(n))

	return math.Round(sqrt) == sqrt, int(sqrt)
}
