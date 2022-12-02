package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	runningCount := 0
	highestTotal := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			if runningCount > highestTotal {
				highestTotal = runningCount
			}

			runningCount = 0
		} else {
			value, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}

			runningCount += value
		}
	}

	fmt.Println("Highest Total: ", highestTotal)
}