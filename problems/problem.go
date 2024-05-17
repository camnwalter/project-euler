package problems

import "fmt"

func Problem(num int) {
	problems := []func(){
		One, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Eleven, Twelve, Thirteen, Fourteen, Fifteen,
		Sixteen, Seventeen, Eighteen, Nineteen,
	}

	if num == 0 {
		fmt.Println("Running all problems.")
		fmt.Println()

		for i, problem := range problems {
			fmt.Print("Problem ", i+1, ": ")
			problem()
		}
	} else {
		fmt.Print("Problem ", num, ": ")
		problems[num-1]()
	}
}
