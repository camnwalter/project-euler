package problems

import "fmt"

func Twelve() {
	n := 1
	num := Triangle(n)
	for NumDivisors(num) < 500 {
		n++
		num = Triangle(n)
	}

	fmt.Println("triangle divisors =", num)
}

func Triangle(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}

func NumDivisors(n int) int {
	if n == 1 {
		return 1
	}

	divisors := 0
	max := n / 2

	for i := 1; i < max; i++ {
		if n%i == 0 {
			divisors += 2
			max = n / i
		}
	}

	return divisors
}
