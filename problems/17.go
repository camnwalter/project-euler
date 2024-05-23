package problems

import "strings"

var small = []string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
	"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
}

var tens = []string{
	"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
}

func Seventeen() int {
	length := 0

	replacer := strings.NewReplacer(" ", "", "-", "")

	for n := 1; n <= 1000; n++ {
		length += len(replacer.Replace(intToWord(n)))
	}

	return length
}

func intToWord(n int) string {
	if n <= 19 {
		return small[n]
	}

	if n <= 99 {
		rest := n % 10

		if rest == 0 {
			return tens[n/10]
		}

		return tens[n/10] + "-" + intToWord(rest)
	}

	if n <= 999 {
		rest := n % 100

		if rest == 0 {
			return small[n/100] + " hundred"
		}

		return small[n/100] + " hundred and " + intToWord(rest)
	}

	return "one thousand"
}
