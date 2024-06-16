package problems

import (
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/camnwalter/project-euler/utils"
)

func Problem(num int) {
	problems := []func() int{
		One, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Eleven,
		Twelve, Thirteen, Fourteen, Fifteen, Sixteen, Seventeen, Eighteen,
		Nineteen, Twenty, TwentyOne, TwentyTwo, TwentyThree, TwentyFour,
		TwentyFive, TwentySix, TwentySeven, TwentyEight, TwentyNine, Thirty,
		ThirtyOne, ThirtyTwo, ThirtyThree, ThirtyFour, ThirtyFive, ThirtySix,
		ThirtySeven, ThirtyEight, ThirtyNine, Forty, FortyOne, FortyTwo,
		FortyThree, FortyFour, FortyFive, FortySix, FortySeven, FortyEight,
		FortyNine, Fifty, FiftyOne, FiftyTwo, FiftyThree, FiftyFour, FiftyFive,
		FiftySix, FiftySeven, FiftyEight, FiftyNine, Sixty, SixtyOne, SixtyTwo,
		SixtyThree, SixtyFour, SixtyFive, SixtySix, SixtySeven, SixtyEight,
		SixtyNine, Seventy, SeventyOne, SeventyTwo, SeventyThree, SeventyFour,
		SeventyFive, SeventySix, SeventySeven, SeventyEight,
	}

	answers, _ := utils.GetFileLines("answers.txt")
	start := time.Now()

	log.SetFlags(0)

	if num == 0 {
		log.Println("Running all problems.")
		log.Println()

		var wg sync.WaitGroup

		for i, problem := range problems {
			wg.Add(1)

			go func(i int, problem func() int) {
				defer wg.Done()
				if problem == nil {
					return
				}

				actual := strconv.Itoa(problem())
				if i < len(answers) && actual != answers[i] {
					log.Printf("Problem %d: Incorrect! Expected %s, got %s\n", i+1, answers[i], actual)
				} else {
					log.Printf("Problem %d: %s\n", i+1, actual)
				}
			}(i, problem)
		}

		wg.Wait()
		log.Printf("Time elapsed: %d ms\n", time.Since(start).Milliseconds())
	} else if num <= len(problems) {
		i := num - 1
		actual := strconv.Itoa(problems[i]())

		if i < len(answers) && answers[i] != "" && actual != answers[i] {
			log.Fatalf("Problem %d: Incorrect! Expected %s, got %s\n", i+1, answers[i], actual)
		} else {
			log.Printf("Problem %d: %s | Time elapsed: %d ms\n", i+1, actual, time.Since(start).Milliseconds())
		}
	} else {
		log.Fatalf("Invalid problem number. Valid numbers are 1-%d\n", len(problems))
	}
}
