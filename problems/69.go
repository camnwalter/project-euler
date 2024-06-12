package problems

func SixtyNine() int {
	return 2 * 3 * 5 * 7 * 11 * 13 * 17
}

func totient(n int, primes []int) int {
	// copied optimized formula

	result := n

	for i := 0; primes[i]*primes[i] <= n; i++ {
		if n%primes[i] == 0 {
			result -= (result / primes[i])

			for n%primes[i] == 0 {
				n /= primes[i]
			}
		}
	}

	if n > 1 {
		result -= result / n
	}

	return result
}
