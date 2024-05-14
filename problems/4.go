package problems

import (
	"fmt"
	"strconv"
)

func Four() {
	max := 0

	a := 100

	for a < 1000 {
		b := 100

		for b < 1000 {
			mult := a * b
			if IsPalindromic(mult) && mult > max {
				max = mult
			}

			b++
		}

		a++
	}

	fmt.Println("max palindrome =", max)
}

func IsPalindromic(n int) bool {
	str := strconv.Itoa(n)

	front := 0
	back := len(str) - 1

	for front <= back {
		if str[front] != str[back] {
			return false
		}

		front++
		back--
	}

	return true
}
