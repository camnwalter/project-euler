package problems

import "github.com/camnwalter/project-euler/big"

func TwentyFive() int {
	n := 0
	cache := make(map[int]*big.MyBigInt)
	for len(bigFib(n, cache).Digits) < 1000 {
		n++
	}

	return n + 1
}

func bigFib(n int, cache map[int]*big.MyBigInt) *big.MyBigInt {
	if data, ok := cache[n]; ok {
		return data
	}

	if n <= 1 {
		cache[n] = big.New(1)
		return cache[n]
	}

	cache[n], _ = bigFib(n-1, cache).Plus(bigFib(n-2, cache))

	return cache[n]
}
