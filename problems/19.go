package problems

var days = []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}

func Nineteen() int {
	sundayFirsts := 0
	index := 1

	for month := 0; month < 12; month++ {
		index += daysInMonth(month, 1900)
		index %= len(days)
	}

	for year := 1901; year <= 2000; year++ {
		for month := 0; month < 12; month++ {
			if index == 0 {
				sundayFirsts++
			}

			index += daysInMonth(month, year)
			index %= len(days)
		}
	}

	return sundayFirsts
}

func daysInMonth(month int, year int) int {
	switch month {
	case 1:
		if year%4 != 0 || (year%100 == 0 && year%400 != 0) {
			return 28
		}
		return 29

	case 3, 5, 8, 10:
		return 30

	default:
		return 31
	}
}
