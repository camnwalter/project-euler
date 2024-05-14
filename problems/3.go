package problems

import (
	"fmt"
)

func Three() {
	num := 600851475143
	max := 1

	for i := 2; max < num; i++ {
		if num%i == 0 {
			max = i
			num /= i
		}
	}

	fmt.Println("max =", max)
}
