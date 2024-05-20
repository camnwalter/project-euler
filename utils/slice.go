package utils

import (
	"fmt"
	"slices"
)

func AddIfAbsent[T comparable](slice []T, values ...T) []T {
	for _, value := range values {
		if !slices.Contains(slice, value) {
			slice = append(slice, value)
		}
	}

	return slice
}

func Swap[T any](slice []T, a int, b int) {
	slice[a], slice[b] = slice[b], slice[a]
}

func Permutations[T any](slice []T) [][]T {
	return generate(len(slice), slice, make([][]T, 0))
}

func generate[T any](n int, A []T, out [][]T) [][]T {
	// c is an encoding of the stack state. c[k] encodes the for-loop counter for when generate(k - 1, A) is called
	c := make([]int, n)

	out = append(out, slices.Clone(A))

	// i acts similarly to a stack pointer
	i := 1
	for i < n {
		if c[i] < i {
			if i%2 == 0 {
				Swap(A, 0, i)
			} else {
				Swap(A, c[i], i)
			}

			out = append(out, slices.Clone(A))
			// Swap has occurred ending the for-loop. Simulate the increment of the for-loop counter
			c[i]++
			// Simulate recursive call reaching the base case by bringing the pointer to the base case analog in the array
			i = 1
		} else {
			// Calling generate(i+1, A) has ended as the for-loop terminated. Reset the state and simulate popping the stack by incrementing the pointer.
			c[i] = 0
			i++
		}
	}

	return out
}

func Join[T any](slice []T, separator string) string {
	out := ""

	for i, v := range slice {
		out += fmt.Sprint(v)

		if i < len(slice) {
			out += fmt.Sprint(separator)
		}
	}
	return out
}

func DefaultSlice[T any](size int, initial T) []T {
	out := make([]T, size)
	for i := range size {
		out[i] = initial
	}

	return out
}

func Filter[T any](slice []T, filterFunc func(val T, index int) bool) []T {
	out := make([]T, 0)

	for i, v := range slice {
		if filterFunc(v, i) {
			out = append(out, v)
		}
	}

	return out
}

func Map[In any, Out any](slice []In, mapFunc func(val In, index int) Out) []Out {
	out := make([]Out, 0)

	for i, v := range slice {
		out = append(out, mapFunc(v, i))
	}

	return out
}
