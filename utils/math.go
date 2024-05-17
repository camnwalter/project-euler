package utils

import (
	"strconv"
	"strings"
)

func Min(nums ...int) int {
	min := nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min
}

func Max(nums ...int) int {
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

func ToDigitArray(n int) []int {
	out := make([]int, 0)

	for _, digit := range strings.Split(Reverse(strconv.Itoa(n)), "") {
		digit, _ := strconv.Atoi(digit)
		out = append(out, digit)
	}

	return out
}

func SumArray(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}
