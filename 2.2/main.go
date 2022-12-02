package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	Run()
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func Run() {
	file, err := os.Open("input.txt")
	handleError(err)

	defer file.Close()

	var totalScore int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		data := strings.Split(input, " ")

		theirShape := 0
		requiredResult := 0
		resultScore := 0

		switch data[0] {
		case "A":
			theirShape = 0
		case "B":
			theirShape = 1
		case "C":
			theirShape = 2
		}

		switch data[1] {
		case "X":
			requiredResult = -1
			resultScore = 0
		case "Y":
			requiredResult = 0
			resultScore = 3
		case "Z":
			requiredResult = 1
			resultScore = 6
		}

		myShape := ((theirShape + 3 + requiredResult) % 3) + 1
	
		totalScore += resultScore
		totalScore += myShape
	}

	fmt.Println("Total score: ", totalScore)
}
