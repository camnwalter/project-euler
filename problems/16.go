package problems

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func Sixteen() {
	bigNumber := big.NewInt(2)
	bigNumber.Exp(bigNumber, big.NewInt(1000), nil)

	digits := bigNumber.String()
	sum := 0
	for _, digit := range strings.Split(digits, "") {
		num, _ := strconv.Atoi(digit)
		sum += num
	}

	fmt.Println("sum =", sum)
}
