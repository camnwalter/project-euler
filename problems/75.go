package problems

import (
	"github.com/camnwalter/project-euler/utils"
)

func SeventyFive() int {
	count := 0
	counts := make(map[int]int)

	const limit int = 1_500_000

	for n := 1; n <= 611; n++ {
		for m := n + 1; m*(m+n) <= 750_000; m++ {
			a := m*m - n*n
			b := 2 * m * n
			c := m*m + n*n

			if utils.GCF(m, n) != 1 {
				continue
			}

			if m%2 == 1 && n%2 == 1 {
				continue
			}

			// 12 <= a+b+c = 2m^2 + 2mn = 2m(m+n) <= 1_500_000
			//
			// 6 <= m(m+n) <= 750_000
			// min m = n+1
			// 6 <= (n+1)(2n+1) <= 750_000
			// 6 <= 2n^2+3n+1 <= 750_000
			//
			// 1 <= n <= 611

			for k := 1; k*(a+b+c) <= limit; k++ {
				counts[k*(a+b+c)]++
			}
		}
	}

	for _, v := range counts {
		if v == 1 {
			count++
		}
	}

	return count
}
