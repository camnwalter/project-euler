package problems

import (
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/camnwalter/project-euler/utils"
)

func SixtyEight() int {
	perms := utils.Permutations([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	max := 0

outer:
	for _, perm := range perms {
		a := []int{perm[5], perm[4], perm[3]}

		total := utils.SumArray(a)

		b := []int{perm[9], perm[3], perm[2]}
		c := []int{perm[8], perm[2], perm[1]}
		d := []int{perm[7], perm[1], perm[0]}
		e := []int{perm[6], perm[0], perm[4]}

		all := [][]int{a, b, c, d, e}

		for _, nums := range all {
			if utils.SumArray(nums) != total {
				continue outer
			}
		}

		minSlice := slices.MinFunc(all, func(a []int, b []int) int {
			return a[0] - b[0]
		})
		startIndex := slices.IndexFunc(all, func(val []int) bool {
			return slices.Equal(val, minSlice)
		})

		out := append(all[startIndex:], all[0:startIndex]...)

		var buf strings.Builder
		for _, nums := range out {
			for _, num := range nums {
				buf.WriteString(strconv.Itoa(num))
			}
		}

		str := buf.String()
		if len(str) == 17 {
			continue
		}

		value, _ := strconv.Atoi(str)
		log.Println(str, total)
		if value > max {
			max = value
		}
	}

	return max
}
