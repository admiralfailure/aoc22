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
		myShape := 0

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
			myShape = 0
		case "Y":
			myShape = 1
		case "Z":
			myShape = 2
		}

		totalScore += myShape + 1

		if ((theirShape + 1) % 3) == myShape {
			totalScore += 6
		} else if theirShape == myShape {
			totalScore += 3
		}
	}

	fmt.Println("Total score: ", totalScore)
}
