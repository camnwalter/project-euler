package problems

func SixtyNine() int {
	return 2 * 3 * 5 * 7 * 11 * 13 * 17
}

func totient(n int, primes []int) int {
	count := float64(n)

	for _, factor := range primeFactorize(n, primes) {
		diff := 1 - 1/float64(factor)
		count *= diff
	}

	return int(count)
}
