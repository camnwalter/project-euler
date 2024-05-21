package problems

import (
	"fmt"

	"github.com/camnwalter/project-euler/utils"
)

func Forty() {
	product := 1
	for n := 1; n <= 1_000_000; n *= 10 {
		product *= getNthDigit(n)
	}
	fmt.Println(product)
}

func getNthDigit(n int) int {
	//					  1, 2,  3,   4,     5,      6,       7
	digitCutoffs := []int{1, 10, 190, 2_890, 38_890, 488_890, 5_888_890}
	// 9 digits  2*90=180 digits 3*900=2700 digits

	for i := 0; i < len(digitCutoffs); i++ {
		spot := digitCutoffs[i]
		nextSpot := digitCutoffs[utils.Min(len(digitCutoffs)-1, i+1)]

		if n >= nextSpot {
			continue
		}

		if n < spot {
			break
		}

		digitsPerNum := i + 1
		num := utils.Pow(10, i)

		var j int
		for j = spot; j < n; j += digitsPerNum {
			num++
		}

		if j == n {
			return num / utils.Pow(10, i)
		} else {
			num--

			goLeft := j - n

			return (num % utils.Pow(10, goLeft)) / utils.Pow(10, goLeft-1)
		}
	}

	return -1
}
