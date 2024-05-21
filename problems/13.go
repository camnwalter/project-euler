package problems

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Thirteen() {
	bytes, err := os.ReadFile("inputs/13.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	nums := strings.Split(string(bytes), "\n")

	carry := 0

	digits := make([]string, 10)

	for col := len(nums[0]) - 1; col >= 0; col-- {
		sum := carry
		for _, num := range nums {
			digit := int(num[col] - '0')
			sum += digit
		}

		if col < 10 {
			digits[col] = strconv.Itoa(sum % 10)
		}
		carry = sum / 10
	}

	if carry > 0 {
		carryString := strconv.Itoa(carry)
		digitsToReplace := len(carryString)

		fmt.Print("digits = ")
		fmt.Print(carry)
		fmt.Println(strings.Join(digits[0:10-digitsToReplace], ""))
	} else {
		fmt.Println("digits =", strings.Join(digits, ""))
	}
}
