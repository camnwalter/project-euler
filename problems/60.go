package problems

import (
	"github.com/camnwalter/project-euler/utils"
)

func Sixty() int {
	primes := utils.PrimeSieve(10_000)
	length := len(primes)

	for i := 0; i < length-4; i++ {
		a := primes[i]
		for j := i + 1; j < length-3; j++ {
			b := primes[j]
			if !validPair(a, b) {
				continue
			}

			for k := j + 1; k < length-2; k++ {
				c := primes[k]
				if !validPair(a, c) || !validPair(b, c) {
					continue
				}

				for l := k + 1; l < length-1; l++ {
					d := primes[l]
					if !validPair(a, d) || !validPair(b, d) || !validPair(c, d) {
						continue
					}

					for m := l + 1; m < length; m++ {
						e := primes[m]
						if validPair(a, e) && validPair(b, e) && validPair(c, e) && validPair(d, e) {
							return a + b + c + d + e
						}
					}
				}
			}
		}
	}

	return -1
}

func validPair(a int, b int) bool {
	a, b = pairNumbers(a, b)

	return utils.IsPrime(a) && utils.IsPrime(b)
}

func pairNumbers(a int, b int) (int, int) {
	aDigits := utils.ToDigitArray(a)
	bDigits := utils.ToDigitArray(b)

	out1 := append(aDigits, bDigits...)
	out2 := append(bDigits, aDigits...)

	return utils.FromDigitArray(out1), utils.FromDigitArray(out2)
}
