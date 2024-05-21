package problems

import (
	"fmt"
	"os"
	"strings"
)

func FortyTwo() {
	bytes, err := os.ReadFile("inputs/42.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	words := strings.Split(strings.ReplaceAll(string(bytes), "\"", ""), ",")

	triangles := 0

	for _, word := range words {
		wordSum := 0
		for _, c := range word {
			wordSum += int(c - 'A' + 1)
		}

		if IsTriangleNumber(wordSum) {
			triangles++
		}
	}

	fmt.Println("number of triangle words =", triangles)
}

func IsTriangleNumber(n int) bool {
	// i(i+1)/2 = n
	// i(i+1) = 2n

	for i := 1; i*i+1 <= 2*n; i++ {
		if i*(i+1)/2 == n {
			return true
		}
	}

	return false
}
