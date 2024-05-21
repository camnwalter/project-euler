package problems

import (
	"fmt"
	"math"
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

func IsTriangleNumber(t int) bool {
	// t = n (n + 1) / 2
	// .5n^2 + .5n - t = 0
	// t = (-1 + sqrt(1 + 8t)) / 2

	n := (-1 + math.Sqrt(float64(8*t+1))) / 2
	return math.Round(n) == n
}
