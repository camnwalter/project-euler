package problems

import (
	"fmt"
)

func Three() {
	num := 600851475143
	max := 1

	for i := 2; max < num; i++ {
		if num%i == 0 {
			max = i
			num /= i
		}
	}

	fmt.Println("max =", max)
}

func IsPrime(n int) bool {
	max := n / 2
	for i := 2; i < max; {
		if n%i == 0 {
			return false
		}

		i++
		max = n / i
	}

	return true
}
