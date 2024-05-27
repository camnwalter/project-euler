package problems

import (
	"os"
	"strconv"
	"strings"

	"github.com/camnwalter/project-euler/utils"
)

func FiftyNine() int {
	bytes, _ := os.ReadFile("inputs/59.txt")
	numbers := utils.Map(strings.Split(string(bytes), ","), func(a string, _ int) byte {
		num, _ := strconv.Atoi(a)
		return byte(num)
	})

	password := []byte{'e', 'x', 'p'}

	sum := 0

	for i := 0; i < len(numbers); i++ {
		sum += int(numbers[i] ^ password[i%3])
	}

	return sum
}

// Used to find the password (append to file and find something that looks word-like)
// func endsWord(b byte) bool {
// 	return b == ' ' || b == '\n' || b == '\r' || b == '.' || b == ',' ||
// 		b == '(' || b == ')' || b == '!' || b == '?'
// }

// func validWord(word []byte) bool {
// 	for _, b := range word {
// 		if !(('A' <= b && b <= 'Z') || ('a' <= b && b <= 'z') || b == '\'') {
// 			return false
// 		}
// 	}

// 	return true
// }
