package problems

import (
	"fmt"
	"slices"

	"github.com/camnwalter/project-euler/utils"
)

func FortyNine() {
	primes := PrimeSieve(10_000)

outer:
	for _, p1 := range primes {
		if p1 == 1487 {
			continue
		}

		for _, p2 := range primes {
			if p2 == p1 || !IsPermutation(p1, p2) {
				continue
			}

			for _, p3 := range primes {
				if p3 == p1 || p3 == p2 || !IsPermutation(p1, p3) {
					continue
				}

				if p1-p2 == p2-p3 {
					fmt.Printf("output = %d%d%d\n", p1, p2, p3)
					break outer
				}
			}
		}
	}
}

func IsPermutation(a int, b int) bool {
	aDigits := utils.ToDigitArray(a)
	slices.Sort(aDigits)
	bDigits := utils.ToDigitArray(b)
	slices.Sort(bDigits)

	if len(aDigits) != len(bDigits) {
		return false
	}

	for i := range len(aDigits) {
		if aDigits[i] != bDigits[i] {
			return false
		}
	}

	return true
}
