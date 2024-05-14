package problems

import "fmt"

func Ten() {
	limit := 2_000_000

	sum := 0
	prime := 2

	for prime < limit {
		sum += prime
		prime = NextPrime(prime)
	}

	fmt.Println("sum of primes =", sum)
}
