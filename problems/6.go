package problems

func Six() int {
	squaredSum := 0
	sumSquared := 0

	for i := range 101 {
		squaredSum += i
		sumSquared += i * i
	}

	squaredSum *= squaredSum

	return squaredSum - sumSquared
}
