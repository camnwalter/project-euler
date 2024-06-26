package problems

import (
	"os"
	"strings"
)

func Eight() int {
	bytes, _ := os.ReadFile("inputs/8.txt")
	contents := strings.ReplaceAll(string(bytes), "\n", "")

	largestProduct := 0

outer:
	for i := 0; i < len(contents)-13; i++ {
		product := 1

		for j := i; j < i+13; j++ {
			num := int(contents[j] - '0')

			if num == 0 {
				i = j
				continue outer
			}

			product *= num
		}

		if product > largestProduct {
			largestProduct = product
		}
	}

	return largestProduct
}
