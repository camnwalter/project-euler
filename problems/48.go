package problems

import "fmt"

func FortyEight() {
	sum := 0
	mod := 10_000_000_000

	// only care about last 10 digits

	for n := 1; n <= 1000; n++ {
		product := 1
		for range n {
			product *= n
			product %= mod
		}

		sum += product
		sum %= mod
	}

	fmt.Println(sum)
}
