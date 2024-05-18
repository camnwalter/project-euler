package problems

import (
	"fmt"
	"slices"
)

func TwentyFive() {
	n := 0
	cache := make(map[int][]int)
	for len(BigFib(n, cache)) < 1000 {
		n++
	}

	fmt.Println("1000 digits of fib index =", n+1)
}

func BigFib(n int, cache map[int][]int) []int {
	if data, ok := cache[n]; ok {
		return data
	}

	if n <= 1 {
		cache[n] = []int{1}
		return cache[n]
	}

	cache[n] = arrayAdd(slices.Clone(BigFib(n-1, cache)), slices.Clone(BigFib(n-2, cache)))

	return cache[n]
}

func arrayAdd(data []int, n []int) []int {
	carry := 0
	j := 0
	for i := 0; i < len(data); i++ {
		addend := 0
		if j < len(n) {
			addend = n[j]
			j++
		}

		sum := data[i] + addend + carry

		data[i] = sum % 10
		carry = sum / 10
	}

	for j < len(n) {
		data = append(data, n[j])
		j++
	}

	for carry != 0 {
		data = append(data, carry%10)
		carry /= 10
	}

	return data
}
