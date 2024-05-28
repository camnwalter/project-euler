package problems

import (
	"slices"
	"strconv"
	"strings"

	"github.com/camnwalter/project-euler/utils"
)

func SixtyTwo() int {
	seen := make(map[string][]int)

	for n := 1; ; n++ {
		cubed := n * n * n

		word := strconv.Itoa(cubed)
		digits := strings.Split(word, "")
		slices.Sort(digits)

		sorted := strings.Join(digits, "")

		if _, ok := seen[sorted]; !ok {
			seen[sorted] = make([]int, 2)
		}

		seen[sorted][0]++
		if seen[sorted][1] == 0 {
			seen[sorted][1] = cubed
		}

		seen[sorted][1] = utils.Min(seen[sorted][1], cubed)

		if seen[sorted][0] == 5 {
			return seen[sorted][1]
		}
	}
}
