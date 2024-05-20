package problems

import (
	"fmt"

	"github.com/camnwalter/project-euler/utils"
)

func Four() {
	max := 0

	a := 100

	for a < 1000 {
		b := 100

		for b < 1000 {
			mult := a * b
			if IsPalindrome(fmt.Sprint(mult)) && mult > max {
				max = mult
			}

			b++
		}

		a++
	}

	fmt.Println("max palindrome =", max)
}

func IsPalindrome(n string) bool {
	return n == utils.Reverse(n)
}
