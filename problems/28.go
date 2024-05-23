package problems

func TwentyEight() int {
	// 3, 13, 31 => +2, +10, +18
	// 5, 17, 37 => +4, +12, +20
	// 7, 21, 43 => +6, +14, +22
	// 9, 25, 49 => +8, +16, +24

	a, b, c, d := 1, 1, 1, 1
	n := 2

	sum := 1

	for step := 1; step*2+1 <= 1001; step++ {
		a += n
		n += 2

		b += n
		n += 2

		c += n
		n += 2

		d += n
		n += 2

		sum += a + b + c + d
	}

	return sum
}
