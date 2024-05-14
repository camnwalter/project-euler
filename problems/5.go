package problems

import "fmt"

func Five() {
	num := 20

	for i := 2; i <= 20; i++ {
		if num%i != 0 {
			num++

			i = 2
		}
	}

	fmt.Println("evenly divisible =", num)
}
