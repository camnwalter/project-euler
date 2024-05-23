package main

import (
	"log"
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
		log.Fatal(err)
	}

	problems.Problem(problemNumber)
}
