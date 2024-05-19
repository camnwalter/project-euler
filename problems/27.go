package problems

import "fmt"

func TwentySeven() {
	var max int
	var maxA int
	var maxB int

	for a := -999; a <= 999; a++ {
		for b := -1000; b <= 1000; b++ {
			n := 0
			for ; IsPrime(n*n + a*n + b); n++ {
			}

			if n > max {
				max = n
				maxA = a
				maxB = b
			}
		}
	}

	fmt.Println("max a*b =", maxA*maxB)
}
