package problems

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/camnwalter/project-euler/utils"
)

func Eleven() {
	bytes, err := os.ReadFile("inputs/11.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	file := string(bytes)

	nums := make([][]int, 20)

	lines := strings.Split(file, "\n")
	for r, line := range lines {

		rowData := strings.Split(line, " ")

		row := make([]int, 20)
		for c, num := range rowData {
			num, _ := strconv.Atoi(num)
			row[c] = num
		}

		nums[r] = row
	}

	maxHorizontal := 0
	maxVertical := 0
	maxDiagonal := 0

	for _, r := range nums {
		product := FindMaxProduct(r, 4)
		maxHorizontal = utils.Max(maxHorizontal, product)
	}

	for _, r := range Transpose(nums) {
		product := FindMaxProduct(r, 4)
		maxVertical = utils.Max(maxVertical, product)
	}

	for r := 0; r < len(nums); r++ {
		for c := 0; c < len(nums[r]); c++ {
			rightDiagonal := make([]int, 4)
			leftDiagonal := make([]int, 4)
			for n := 0; n < 4; n++ {
				if r+n < len(nums) && c+n < len(nums[r]) {
					rightDiagonal[n] = nums[r+n][c+n]
				}

				if r+n < len(nums) && c-n >= 0 {
					leftDiagonal[n] = nums[r+n][c-n]
				}
			}

			productRight := FindMaxProduct(rightDiagonal, 4)
			productLeft := FindMaxProduct(leftDiagonal, 4)

			maxDiagonal = utils.Max(maxDiagonal, productRight, productLeft)
		}
	}

	fmt.Println("max of four in each direction =", utils.Max(maxHorizontal, maxVertical, maxDiagonal))
}

func FindMaxProduct(data []int, terms int) int {
	max := 0

	for i := 0; i < len(data); i++ {
		product := 1
		for j := i; j < i+terms && j < len(data); j++ {
			product *= data[j]
		}

		max = utils.Max(max, product)
	}

	return max
}

func Transpose(data [][]int) [][]int {
	transposed := make([][]int, len(data))

	for r := 0; r < len(data); r++ {
		transposed[r] = make([]int, len(data[r]))
		for c := 0; c < len(data[r]); c++ {
			transposed[r][c] = data[c][r]
		}
	}

	return transposed
}
