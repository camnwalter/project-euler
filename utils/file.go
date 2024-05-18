package utils

import (
	"os"
	"strings"
)

func GetFileLines(path string) ([]string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(bytes), "\n"), nil
}
