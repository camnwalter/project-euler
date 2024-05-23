package problems

import (
	"fmt"

	"github.com/camnwalter/project-euler/utils"
)

func Four() int {
	max := 0

	a := 100

	for a < 1000 {
		b := 100

		for b < 1000 {
			mult := a * b
			if utils.IsPalindrome(fmt.Sprint(mult)) && mult > max {
				max = mult
			}

			b++
		}

		a++
	}

	return max
}
