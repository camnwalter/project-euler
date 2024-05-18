package problems

import (
	"fmt"
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

	cache[n] = arrayAdd(BigFib(n-1, cache), BigFib(n-2, cache))

	return cache[n]
}

func arrayAdd(a []int, b []int) []int {
	out := make([]int, 0)

	carry := 0
	for i := 0; i < len(a); i++ {
		addend := 0
		if i < len(b) {
			addend = b[i]
		}

		sum := a[i] + addend + carry

		out = append(out, sum%10)
		carry = sum / 10
	}

	for i := len(a); i < len(b); i++ {
		sum := b[i] + carry

		out = append(out, sum%10)
		carry = sum / 10
	}

	for carry != 0 {
		out = append(out, carry%10)
		carry /= 10
	}

	return out
}
