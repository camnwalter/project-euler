package problems

import (
	"strconv"
	"strings"

	"github.com/camnwalter/project-euler/utils"
)

func EightyOne() int {
	lines, _ := utils.GetFileLines("inputs/81.txt")

	matrix := make([][]int, len(lines))

	for i, line := range lines {
		matrix[i] = make([]int, len(lines))

		for j, v := range strings.Split(line, ",") {
			matrix[i][j], _ = strconv.Atoi(v)
		}
	}

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			up := 0
			hasUp := false
			if i > 0 {
				up = matrix[i-1][j]
				hasUp = true
			}

			left := 0
			hasLeft := false
			if j > 0 {
				left = matrix[i][j-1]
				hasLeft = true
			}

			if hasLeft {
				if hasUp {
					matrix[i][j] += utils.Min(up, left)
				} else {
					matrix[i][j] += left
				}
			} else {
				if hasUp {
					matrix[i][j] += up
				}
			}

		}
	}

	return matrix[len(matrix)-1][len(matrix)-1]
}
