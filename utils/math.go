package utils

import (
	"slices"
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

func FromDigitArray(n []int) int {
	out := 0

	for i, digit := range n {
		out += digit * Pow(10, i)
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

func Pow(a int, b int) int {
	product := 1

	for range b {
		product *= a
	}

	return product
}

func GetFactors(n int) []int {
	out := make([]int, 0)

	out = AddIfAbsent(out, 1)
	out = AddIfAbsent(out, n)

	if n == 1 {
		return out
	}

	max := n / 2

	for i := 2; i <= max; i++ {
		if n%i == 0 {
			max = n / i
			out = AddIfAbsent(out, i, max)
		}
	}

	return out
}

func GCF(a int, b int) int {
	aFactors := GetFactors(a)
	bFactors := GetFactors(b)

	commonFactors := make([]int, 0)
	for _, factor := range aFactors {
		if slices.Contains(bFactors, factor) {
			commonFactors = AddIfAbsent(commonFactors, factor)
		}
	}

	for _, factor := range bFactors {
		if slices.Contains(aFactors, factor) {
			commonFactors = AddIfAbsent(commonFactors, factor)
		}
	}

	return Max(commonFactors...)
}

// place (ones = 1, tens = 2, etc)
func GetDigit(n int, digit int) int {
	return (n % Pow(10, digit)) / Pow(10, digit-1)
}

// place (ones = 1, tens = 2, etc)
func GetDigits(n int, min int, max int) int {
	return (n % Pow(10, max)) / Pow(10, min-1)
}
