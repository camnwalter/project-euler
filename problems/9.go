package problems

import "fmt"

func Nine() {
	var a int
	var b int
	var c int

outer:
	for a = 1; a < 1000; a++ {
		for b = 1; b < 1000; b++ {
			for c = 1; c < 1000; c++ {
				if a*a+b*b == c*c && a+b+c == 1000 {
					break outer
				}
			}
		}
	}

	fmt.Println("abc =", a*b*c)
}
