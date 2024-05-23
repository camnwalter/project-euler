package problems

func Three() int {
	num := 600851475143
	max := 1

	for i := 2; max < num; i++ {
		if num%i == 0 {
			max = i
			num /= i
		}
	}

	return max
}
