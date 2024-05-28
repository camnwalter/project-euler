package problems

import "github.com/camnwalter/project-euler/utils"

func Twelve() int {
	n := 1
	num := utils.Triangle(n)
	for len(utils.GetProperDivisors(num))+1 < 500 {
		n++
		num = utils.Triangle(n)
	}

	return num
}
