package utils

import "slices"

func AddIfAbsent[val comparable](slice []val, values ...val) []val {
	for _, value := range values {
		if !slices.Contains(slice, value) {
			slice = append(slice, value)
		}
	}

	return slice
}
