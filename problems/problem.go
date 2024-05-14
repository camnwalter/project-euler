package problems

import "fmt"

func Problem(num int) {
	problems := []func(){One, Two, Three, Four}

	if num == 0 {
		fmt.Println("Running all problems.")
		fmt.Println()

		for i, problem := range problems {
			fmt.Printf("Problem %d:\n", i+1)
			problem()
		}
	} else {
		fmt.Printf("Problem %d:\n", num)
		problems[num-1]()
	}
}
