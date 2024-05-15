package problems

import (
	"fmt"
)

func Fifteen() {
	fmt.Println(Combination(40, 20))
}

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}

	product := 1

	for i := 1; i <= n; i++ {
		product *= i
	}

	return product
}

func Combination(n int, k int) int {
	product := 1
	for i := 0; i < k; i++ {
		product *= (n - i)
		product /= (i + 1)
	}
	return product
}
