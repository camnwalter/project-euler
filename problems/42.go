package problems

import (
	"os"
	"strings"

	"github.com/camnwalter/project-euler/utils"
)

func FortyTwo() int {
	bytes, _ := os.ReadFile("inputs/42.txt")

	words := strings.Split(strings.ReplaceAll(string(bytes), "\"", ""), ",")

	triangles := 0

	for _, word := range words {
		wordSum := 0
		for _, c := range word {
			wordSum += int(c - 'A' + 1)
		}

		if utils.IsTriangle(wordSum) {
			triangles++
		}
	}

	return triangles
}
