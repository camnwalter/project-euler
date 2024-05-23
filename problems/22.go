package problems

import (
	"log"
	"slices"
	"strings"

	"github.com/camnwalter/project-euler/utils"
)

func TwentyTwo() int {
	lines, err := utils.GetFileLines("inputs/22.txt")
	if err != nil {
		log.Fatal(err)
	}

	names := strings.Split(lines[0], ",")
	slices.Sort(names)

	sum := 0
	for i, name := range names {
		sum += (i + 1) * nameValue(name)
	}

	return sum
}

func nameValue(name string) int {
	sum := 0

	for _, r := range name {
		if r == '"' {
			continue
		}

		sum += int(r-'A') + 1
	}
	return sum
}
