package problems

import "github.com/camnwalter/project-euler/utils"

func ThirtyNine() int {
	maxSolutions := 0
	maxP := 0

	for p := 1; p <= 1000; p++ {
		solutions := 0

		for a := 1; a < p; a++ {
			for b := 1; b < p-a; b++ {
				if yes, c := utils.IsPerfectSquare(a*a + b*b); yes && a+b+c == p {
					solutions++
				}
			}
		}

		if solutions > maxSolutions {
			maxSolutions = solutions
			maxP = p
		}
	}

	return maxP
}
