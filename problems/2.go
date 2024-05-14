package problems

import "fmt"

func Two() {
	sum := 0
	fib := 1

	for n := 0; fib < 4_000_000; n++ {
		fib = Fib(n)
		if fib%2 == 0 {
			sum += fib
		}
	}

	fmt.Println("sum =", sum)
}

func Fib(n int) int {
	if n <= 1 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}
