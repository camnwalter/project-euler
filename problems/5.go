package problems

func Five() int {
	num := 20

	for i := 2; i <= 20; i++ {
		if num%i != 0 {
			num++

			i = 2
		}
	}

	return num
}
