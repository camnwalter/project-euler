package problems

import "fmt"

func Six() {
	squaredSum := 0
	sumSquared := 0

	for i := range 101 {
		squaredSum += i
		sumSquared += i * i
	}

	squaredSum *= squaredSum

	fmt.Println("difference =", squaredSum-sumSquared)
}
