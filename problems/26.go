package problems

func TwentySix() int {
	max := 0
	maxNum := 0
	for i := 1; i < 1000; i++ {
		_, cur := longDivision(1, i)
		if cur > max {
			max = cur
			maxNum = i
		}
	}

	return maxNum
}

// a / b
//
// returns ([whole, ...decimal], length)
func longDivision(a int, b int) ([]int, int) {
	result := make([]int, 0)

	if a < b {
		a *= 10
		result = append(result, 0)
	}

	seen := make(map[int]int)
	index := 0

	for {
		index++

		if _, ok := seen[a]; ok {
			break
		}

		seen[a] = index

		result = append(result, a/b)

		remainder := a % b
		if remainder == 0 {
			return result, 0
		}

		a = remainder * 10
	}

	return result, index - seen[a]
}
