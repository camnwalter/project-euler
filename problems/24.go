package problems

import (
	"fmt"
	"slices"

	"github.com/camnwalter/project-euler/utils"
)

func TwentyFour() {
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

	fmt.Println(utils.Join(perms[999_999], ""))
}
