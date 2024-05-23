package problems

import (
	"slices"

	"github.com/camnwalter/project-euler/utils"
)

func TwentyFour() int {
	perms := utils.Permutations([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	slices.SortFunc(perms, func(a []int, b []int) int {
		for i := 0; i < len(a); i++ {
			if a[i] < b[i] {
				return -1
			} else if a[i] > b[i] {
				return 1
			}
		}

		return 0
	})

	slices.Reverse(perms[999_999])

	sum := 0
	place := 1
	for _, digit := range perms[999_999] {
		sum += digit * place
		place *= 10
	}

	return sum
}
