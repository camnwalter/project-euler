package problems

func SixtyThree() int {
	count := 0

	// 1^1, 2^1, 3^1, 4^1, 5^1, 6^1, 7^1, 8^1, 9^1
	count += 9
	// 4^2, 5^2, 6^2, 7^2, 8^2, 9^2
	count += 6
	// 5^3, 6^3, 7^3, 8^3, 9^3
	count += 5
	// 6^4, 7^4, 8^4, 9^4
	count += 4
	// 7^5, 8^5, 9^5
	count += 3
	// 7^6, 8^6, 9^6
	count += 3
	// 8^7, 9^7
	count += 2
	// 8^8, 9^8
	count += 2
	// 8^9, 9^9
	count += 2
	// 8^10, 9^10
	count += 2
	// 9^11
	// ...
	// 9^21
	count += 11

	return count
}
