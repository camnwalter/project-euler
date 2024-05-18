package problems

import (
	"fmt"
	"slices"

	"github.com/camnwalter/project-euler/utils"
)

func TwentyOne() {
	tracked := make([]int, 0)

	for a := 1; a < 10_000; a++ {
		b := d(a)

		if a != b && b < 10_000 && d(b) == a && !slices.Contains(tracked, a) && !slices.Contains(tracked, b) {
			tracked = append(tracked, a, b)
		}
	}

	fmt.Println("sum =", utils.SumArray(tracked))
}

func d(n int) int {
	return utils.SumArray(GetDivisors(n))
}

// get divisors less than n
func GetDivisors(n int) []int {
	out := make([]int, 0)

	out = append(out, 1)

	max := n / 2

	for i := 2; i < max; i++ {
		if n%i == 0 {
			max = n / i
			out = append(out, i, max)
		}
	}

	return out
}