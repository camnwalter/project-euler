package problems

import (
	"math/big"

	"github.com/camnwalter/project-euler/utils"
)

func Eighty() int {
	sum := 0
	for n := 1; n <= 100; n++ {
		if utils.IsSquare(n) {
			continue
		}

		sum += squareRootDigitsSum(n)
	}

	return sum
}

func squareRootDigitsSum(n int) int {
	x := big.NewFloat(float64(n))

	var s big.Float
	s.SetPrec(5000)
	s.Sqrt(x)

	digits := s.Text('f', 101)
	sum := 0

	count := 0
	for i := 0; count < 100; i++ {
		if i == 1 {
			continue
		}

		count++

		sum += int(digits[i] - '0')
	}

	return sum
}
