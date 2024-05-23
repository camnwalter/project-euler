package problems

import "github.com/camnwalter/project-euler/utils"

func Two() int {
	sum := 0
	fib := 1

	for n := 0; fib < 4_000_000; n++ {
		fib = utils.Fib(n)
		if fib%2 == 0 {
			sum += fib
		}
	}

	return sum
}
