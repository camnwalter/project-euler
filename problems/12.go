package problems

import "github.com/camnwalter/project-euler/utils"

func Twelve() int {
	n := 1
	num := triangle(n)
	for len(utils.GetProperDivisors(num))+1 < 500 {
		n++
		num = triangle(n)
	}

	return num
}

func triangle(n int) int {
	return n * (n + 1) / 2
}
