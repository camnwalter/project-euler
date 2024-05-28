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

func ContainsDuplicates[T comparable](slice []T) bool {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] == slice[j] {
				return true
			}
		}
	}

	return false
}

func DefaultSlice[T any](size int, initial T) []T {
	out := make([]T, size)
	for i := range size {
		out[i] = initial
	}

	return out
}

func Every[T any](slice []T, compareFunc func(val T, index int) bool) bool {
	for i, v := range slice {
		if !compareFunc(v, i) {
			return false
		}
	}

	return true
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

func Map[In any, Out any](slice []In, mapFunc func(val In, index int) Out) []Out {
	out := make([]Out, 0)

	for i, v := range slice {
		out = append(out, mapFunc(v, i))
	}

	return out
}

func Permutations[T any](slice []T) [][]T {
	n := len(slice)
	out := make([][]T, 0)

	// c is an encoding of the stack state. c[k] encodes the for-loop counter for when generate(k - 1, slice) is called
	c := make([]int, n)

	out = append(out, slices.Clone(slice))

	// i acts similarly to a stack pointer
	i := 1
	for i < n {
		if c[i] < i {
			if i%2 == 0 {
				Swap(slice, 0, i)
			} else {
				Swap(slice, c[i], i)
			}

			out = append(out, slices.Clone(slice))
			// Swap has occurred ending the for-loop. Simulate the increment of the for-loop counter
			c[i]++
			// Simulate recursive call reaching the base case by bringing the pointer to the base case analog in the array
			i = 1
		} else {
			// Calling generate(i+1, slice) has ended as the for-loop terminated. Reset the state and simulate popping the stack by incrementing the pointer.
			c[i] = 0
			i++
		}
	}

	return out
}

func Swap[T any](slice []T, a int, b int) {
	slice[a], slice[b] = slice[b], slice[a]
}
