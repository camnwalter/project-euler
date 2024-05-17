package problems

import "fmt"

var days = []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}

func Nineteen() {
	sundayFirsts := 0
	index := 1

	for month := 0; month < 12; month++ {
		index += DaysInMonth(month, 1900)
		index %= len(days)
	}

	for year := 1901; year <= 2000; year++ {
		for month := 0; month < 12; month++ {
			if index == 0 {
				sundayFirsts++
			}

			index += DaysInMonth(month, year)
			index %= len(days)
		}
	}

	fmt.Println("1sts that were sundays =", sundayFirsts)
}

func DaysInMonth(month int, year int) int {
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
