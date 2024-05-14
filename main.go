package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/camnwalter/project-euler/problems"
)

func main() {
	if len(os.Args) == 1 {
		problems.Problem(0)
		return
	}

	problemNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Unknown problem number provided")
		return
	}

	problems.Problem(problemNumber)
}
