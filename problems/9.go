package problems

import (
	"math"

	"github.com/camnwalter/project-euler/utils"
)

func Nine() int {
	var a int
	var b int
	var c int

outer:
	for a = 1; a < 1000; a++ {
		for b = 1; b < 1000; b++ {
			if utils.IsSquare(a*a + b*b) {
				cout := int(math.Sqrt(float64(a*a + b*b)))
				if a+b+cout == 1000 {
					c = cout
					break outer
				}
			}
		}
	}

	return a * b * c
}
