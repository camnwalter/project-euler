package problems

func Fourteen() int {
	longestChain := 0
	longestNumber := 0

	for n := 1; n < 1_000_000; n++ {
		collatz := collatz(n)
		if collatz > longestChain {
			longestChain = collatz
			longestNumber = n
		}
	}

	return longestNumber
}

func collatz(n int) int {
	count := 1
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}

		count++
	}

	return count
}
