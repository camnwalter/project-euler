package utils

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func Combination(n int, k int) int {
	product := 1

	if n-k < k {
		k = n - k
	}

	for i := 0; i < k; i++ {
		product *= (n - i)
		product /= (i + 1)
	}
	return product
}

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}

	product := 1

	for i := 1; i <= n; i++ {
		product *= i
	}

	return product
}

func Fib(n int) int {
	if n <= 1 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}

func FromDigitArray(n []int) int {
	out := 0

	for i, digit := range n {
		out += digit * Pow(10, i)
	}

	return out
}

func GCF(a int, b int) int {
	aFactors := GetFactors(a)
	bFactors := GetFactors(b)

	commonFactors := make([]int, 0)
	for _, factor := range aFactors {
		if slices.Contains(bFactors, factor) {
			commonFactors = AddIfAbsent(commonFactors, factor)
		}
	}

	for _, factor := range bFactors {
		if slices.Contains(aFactors, factor) {
			commonFactors = AddIfAbsent(commonFactors, factor)
		}
	}

	return Max(commonFactors...)
}

// place (ones = 1, tens = 2, etc)
func GetDigit(n int, digit int) int {
	return (n % Pow(10, digit)) / Pow(10, digit-1)
}

// place (ones = 1, tens = 2, etc)
func GetDigits(n int, min int, max int) int {
	return (n % Pow(10, max)) / Pow(10, min-1)
}

func GetFactors(n int) []int {
	out := make([]int, 0)

	out = AddIfAbsent(out, 1)
	out = AddIfAbsent(out, n)

	if n == 1 {
		return out
	}

	max := n / 2

	for i := 2; i <= max; i++ {
		if n%i == 0 {
			max = n / i
			out = AddIfAbsent(out, i, max)
		}
	}

	return out
}

func GetProperDivisors(n int) []int {
	out := make([]int, 0)

	if n == 1 {
		return out
	}

	out = append(out, 1)

	max := n / 2

	for i := 2; i <= max; i++ {
		if n%i == 0 {
			max = n / i
			out = AddIfAbsent(out, i, max)
		}
	}

	return out
}

func IsHexagonal(h int) bool {
	// h = n (2n - 1)
	// 2n^2 - n - h = 0
	// n = (1 + sqrt(1 + 8h)) / 4
	n := (1 + math.Sqrt(float64(1+8*h))) / 4
	return math.Round(n) == n
}

func IsHeptagonal(h int) bool {
	// h = n (5n - 3) / 2
	// 2h = n (5n - 3)
	// 5n^2 - 3n - 2h = 0
	// n = (3 + sqrt(9 + 40h)) / 10
	n := (3 + math.Sqrt(float64(9+40*h))) / 10
	return math.Round(n) == n
}

func IsOctagonal(o int) bool {
	// o = n (3n - 2)
	// 3n^2 - 2n - o = 0
	// n = (2 + sqrt(4 + 12o)) / 6
	n := (2 + math.Sqrt(float64(4+12*o))) / 6
	return math.Round(n) == n
}

func IsPentagonal(p int) bool {
	// p = n (3n - 1) / 2
	// 2p = n (3n - 1)
	// 3n^2 - n - 2p = 0
	// n = (1 + sqrt(1 + 24p)) / 6

	n := (1 + math.Sqrt(float64(1+24*p))) / 6
	return math.Round(n) == n
}

func IsPermutation(a int, b int) bool {
	aDigits := ToDigitArray(a)
	slices.Sort(aDigits)
	bDigits := ToDigitArray(b)
	slices.Sort(bDigits)

	if len(aDigits) != len(bDigits) {
		return false
	}

	for i := range len(aDigits) {
		if aDigits[i] != bDigits[i] {
			return false
		}
	}

	return true
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	if n == 2 || n == 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func IsSquare(n int) bool {
	sqrt := math.Sqrt(float64(n))
	return math.Round(sqrt) == sqrt
}

func IsTriangle(t int) bool {
	// t = n (n + 1) / 2
	// 2t = n (n + 1)
	// n^2 + n - 2t = 0
	// t = (-1 + sqrt(1 + 8t)) / 2

	n := (-1 + math.Sqrt(float64(1+8*t))) / 2
	return math.Round(n) == n
}

func Max(nums ...int) int {
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

func Min(nums ...int) int {
	min := nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min
}

func Pow(a int, b int) int {
	product := 1

	for range b {
		product *= a
	}

	return product
}

func PrimeSieve(upperBound int) []int {
	primes := DefaultSlice(upperBound, true)
	primes[0] = false
	primes[1] = false

	for i := 2; i*i < upperBound; i++ {
		if primes[i] {
			for j := i * i; j < upperBound; j += i {
				primes[j] = false
			}
		}
	}

	out := make([]int, 0)
	for num, isPrime := range primes {
		if isPrime {
			out = append(out, num)
		}
	}

	return out
}

func SumArray(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func ToDigitArray(n int) []int {
	out := make([]int, 0)

	for _, digit := range strings.Split(Reverse(strconv.Itoa(n)), "") {
		digit, _ := strconv.Atoi(digit)
		out = append(out, digit)
	}

	return out
}

func Triangle(n int) int {
	return n * (n + 1) / 2
}

func Square(n int) int {
	return n * n
}

func Pentagon(n int) int {
	return n * (3*n - 1) / 2
}

func Hexagon(n int) int {
	return n * (2*n - 1)
}

func Heptagon(n int) int {
	return n * (5*n - 3) / 2
}

func Octagon(n int) int {
	return n * (3*n - 2)
}
