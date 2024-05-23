package problems

import "github.com/camnwalter/project-euler/utils"

func TwentySeven() int {
	var max int
	var maxA int
	var maxB int

	for a := -999; a <= 999; a++ {
		for b := -1000; b <= 1000; b++ {
			n := 0
			for ; utils.IsPrime(n*n + a*n + b); n++ {
			}

			if n > max {
				max = n
				maxA = a
				maxB = b
			}
		}
	}

	return maxA * maxB
}
