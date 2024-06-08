package problems

import (
	"github.com/camnwalter/project-euler/utils"
)

func SixtyOne() int {
	// aabb => bbcc => ccdd => ddee => eeff => ffaa

	for a := 10; a <= 99; a++ {
		for b := 10; b <= 99; b++ {
			aabb := a*100 + b

			if !validFigurate(aabb) {
				continue
			}

			for c := 10; c <= 99; c++ {
				bbcc := b*100 + c
				if !validFigurate(bbcc) {
					continue
				}

				if bbcc == aabb {
					continue
				}

				for d := 10; d <= 99; d++ {
					ccdd := c*100 + d
					if !validFigurate(ccdd) {
						continue
					}

					if ccdd == aabb || ccdd == bbcc {
						continue
					}

					for e := 10; e <= 99; e++ {
						ddee := d*100 + e
						if !validFigurate(ddee) {
							continue
						}

						if ddee == aabb || ddee == bbcc || ddee == ccdd {
							continue
						}

						for f := 10; f <= 99; f++ {
							eeff := e*100 + f

							if !validFigurate(eeff) {
								continue
							}

							if eeff == aabb || eeff == bbcc || eeff == ccdd || eeff == ddee {
								continue
							}

							ffaa := f*100 + a

							if !validFigurate(ffaa) {
								continue
							}

							if ffaa == aabb || ffaa == bbcc || ffaa == ccdd || ffaa == ddee || ffaa == eeff {
								continue
							}

							if getFigurate(aabb, bbcc, ccdd, ddee, eeff, ffaa) {
								return aabb + bbcc + ccdd + ddee + eeff + ffaa
							}
						}
					}
				}
			}
		}
	}

	return 0
}

func validFigurate(n int) bool {
	return utils.IsTriangle(n) || utils.IsSquare(n) || utils.IsPentagonal(n) ||
		utils.IsHexagonal(n) || utils.IsHeptagonal(n) || utils.IsOctagonal(n)
}

func getFigurate(args ...int) bool {
	triangles := make([]int, 0)
	squares := make([]int, 0)
	pentagons := make([]int, 0)
	hexagons := make([]int, 0)
	heptagons := make([]int, 0)
	octagons := make([]int, 0)

	for _, arg := range args {
		if utils.IsTriangle(arg) {
			triangles = append(triangles, arg)
		}

		if utils.IsSquare(arg) {
			squares = append(squares, arg)
		}

		if utils.IsPentagonal(arg) {
			pentagons = append(pentagons, arg)
		}

		if utils.IsHexagonal(arg) {
			hexagons = append(hexagons, arg)
		}

		if utils.IsHeptagonal(arg) {
			heptagons = append(heptagons, arg)
		}

		if utils.IsOctagonal(arg) {
			octagons = append(octagons, arg)
		}
	}

	shapes := [][]int{triangles, squares, pentagons, hexagons, heptagons, octagons}

	for _, shape := range shapes {
		if len(shape) == 0 {
			return false
		}
	}

	for i := 0; i < len(shapes); i++ {
		for j := i + 1; j < len(shapes); j++ {
			shape1 := shapes[i]
			shape2 := shapes[j]

			if len(shape1) == 1 && len(shape2) == 1 {
				if shape1[0] == shape2[0] {
					return false
				}
			}
		}
	}

	return true
}
