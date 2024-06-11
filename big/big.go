package big

import (
	"errors"
	"slices"
	"strconv"
	"strings"
)

const (
	Negative = -1
	Zero     = 0
	Positive = 1
)

type MyBigInt struct {
	Digits []int
	Signum int
}

func New(data int) *MyBigInt {
	Digits := make([]int, 0)
	var signum int

	if data == 0 {
		Digits = append(Digits, 0)
		signum = Zero
	} else {
		if data > 0 {
			signum = Positive
		} else {
			data = -data
			signum = Negative
		}

		for data != 0 {
			Digits = append(Digits, data%10)
			data /= 10
		}
	}

	return &MyBigInt{Digits: Digits, Signum: signum}
}

func (a *MyBigInt) Plus(b *MyBigInt) (*MyBigInt, error) {
	if a.Signum == Negative || b.Signum == Negative {
		return nil, errors.New("negative numbers not supported")
	}

	// 0 + b
	if a.Signum == Zero {
		return &MyBigInt{Digits: slices.Clone(b.Digits), Signum: b.Signum}, nil
	}

	// a + 0
	if b.Signum == Zero {
		return &MyBigInt{Digits: slices.Clone(a.Digits), Signum: a.Signum}, nil
	}

	out := make([]int, 0)
	carry := 0

	for i := 0; i < len(a.Digits); i++ {
		var other int
		if i >= len(b.Digits) {
			other = 0
		} else {
			other = b.Digits[i]
		}

		sum := a.Digits[i] + other + carry
		out = append(out, sum%10)
		carry = sum / 10
	}

	for i := len(a.Digits); i < len(b.Digits); i++ {
		sum := b.Digits[i] + carry
		out = append(out, sum%10)
		carry = sum / 10
	}

	for carry != 0 {
		out = append(out, carry%10)
		carry /= 10
	}

	return &MyBigInt{Digits: out, Signum: Positive}, nil
}

func (a *MyBigInt) Times(b *MyBigInt) *MyBigInt {
	if a.Signum == Zero || b.Signum == Zero {
		return &MyBigInt{Digits: []int{0}, Signum: Zero}
	}

	out := make([]int, len(a.Digits))

	// 12 * 345 = 2*5 + 2*40 + 2*300 + 10*5 + 10*40 + 10*300 = 4140

	for i := 0; i < len(a.Digits); i++ {
		for j := 0; j < len(b.Digits); j++ {
			aDigit := a.Digits[i]
			bDigit := b.Digits[j]

			product := aDigit * bDigit

			index := i + j
			ran := false
			for product != 0 || !ran {
				ran = true

				if index < len(out) {
					out[index] += product % 10
					for out[index] >= 10 {
						out[index] -= 10
						product += 10
					}
				} else {
					out = append(out, product%10)
				}

				product /= 10
				index++
			}
		}
	}

	var signum int
	if a.Signum == Negative {
		if b.Signum == Negative {
			signum = Positive
		} else {
			signum = Negative
		}
	} else {
		if b.Signum == Negative {
			signum = Negative
		} else {
			signum = Positive
		}
	}

	return &MyBigInt{Digits: out, Signum: signum}
}

func Compare(a *MyBigInt, b *MyBigInt) int {
	if a.Signum == Zero && b.Signum == Zero {
		return 0
	}

	if a.Signum < b.Signum {
		return -1
	}

	if a.Signum > b.Signum {
		return 1
	}

	if len(a.Digits) > len(b.Digits) {
		if a.Signum == Negative {
			return -1
		}

		return 1
	}

	if len(a.Digits) < len(b.Digits) {
		if a.Signum == Negative {
			return 1
		}

		return -1
	}

	for i := len(a.Digits) - 1; i >= 0; i-- {
		if a.Digits[i] < b.Digits[i] {
			return -1
		}

		if a.Digits[i] > b.Digits[i] {
			return 1
		}
	}

	return 0
}

func (big *MyBigInt) String() string {
	var buf strings.Builder

	if big.Signum == Negative {
		buf.WriteString("-")
	}

	for i := len(big.Digits) - 1; i >= 0; i-- {
		buf.WriteString(strconv.Itoa(big.Digits[i]))
	}

	return buf.String()
}
