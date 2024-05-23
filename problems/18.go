package problems

import (
	"strconv"
	"strings"

	"github.com/camnwalter/project-euler/utils"
)

func Eighteen() int {
	lines, _ := utils.GetFileLines("inputs/18.txt")

	data := make([][]int, len(lines))

	for i, line := range lines {
		data[i] = make([]int, i+1)

		nums := strings.Split(line, " ")
		for j, num := range nums {
			num, _ := strconv.Atoi(num)
			data[i][j] = num
		}
	}

	return triangleWeight(data)
}

func triangleWeight(data [][]int) int {
	return triangleWeightHelper(data, len(data)-1, 0, 0)
}

// can only go down or down+right
func triangleWeightHelper(data [][]int, targetRow int, startRow int, startCol int) int {
	if startRow >= len(data) || startCol >= len(data[startRow]) {
		return 0
	}

	if targetRow == startRow {
		return data[targetRow][startCol]
	}

	moveLeft := triangleWeightHelper(data, targetRow, startRow+1, startCol)
	moveRight := triangleWeightHelper(data, targetRow, startRow+1, startCol+1)

	return data[startRow][startCol] + utils.Max(moveLeft, moveRight)
}
