package problems

import "fmt"

func Problem(num int) {
	problems := []func(){
		One, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Eleven,
		Twelve, Thirteen, Fourteen, Fifteen, Sixteen, Seventeen, Eighteen,
		Nineteen, Twenty, TwentyOne, TwentyTwo, TwentyThree, TwentyFour,
		TwentyFive, TwentySix, TwentySeven, TwentyEight, TwentyNine, Thirty,
		ThirtyOne, ThirtyTwo, ThirtyThree, ThirtyFour, ThirtyFive, ThirtySix,
		ThirtySeven, ThirtyEight, ThirtyNine, Forty, FortyOne, FortyTwo,
		FortyThree, FortyFour, FortyFive, FortySix, FortySeven, FortyEight,
		FortyNine, Fifty,
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
