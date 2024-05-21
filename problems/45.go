package problems

import (
	"fmt"
	"math"
)

func FortyFive() {
	var n int
	for n = 40_756; ; n++ {
		if IsTriangle(n) && IsPentagonal(n) && IsHexagonal(n) {
			break
		}
	}

	fmt.Println("next num =", n)
}

func IsHexagonal(h int) bool {
	// h = n (2n - 1)
	// 2n^2 - n - h = 0
	// n = (1 + sqrt(1 + 8h)) / 4
	n := (1 + math.Sqrt(float64(1+8*h))) / 4
	return math.Round(n) == n
}
