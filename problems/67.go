package problems

import (
	"strconv"
	"strings"

	"github.com/camnwalter/project-euler/utils"
)

func SixtySeven() int {
	lines, _ := utils.GetFileLines("inputs/67.txt")

	data := make([][]int, len(lines))

	for i, line := range lines {
		data[i] = make([]int, i+1)

		nums := strings.Split(line, " ")
		for j, num := range nums {
			num, _ := strconv.Atoi(num)
			data[i][j] = num
		}
	}

	return triangleWeightAdvanced(data)
}

func triangleWeightAdvanced(data [][]int) int {
	for i, row := range data {
		if i == 0 {
			continue
		}

		for j := range row {
			prevRow := data[i-1]
			left := 0
			if j-1 >= 0 {
				left = prevRow[j-1]
			}

			right := 0
			if j < len(prevRow) {
				right = prevRow[j]
			}
			data[i][j] += utils.Max(left, right)
		}
	}

	return utils.Max(data[len(data)-1]...)
}
