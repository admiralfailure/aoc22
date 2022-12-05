package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Run()

	// s:= "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
	// for _, r := range s {
	// 	fmt.Println(convertRune(r))
	// }
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

	totalPriority := int32(0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split line in half
		line := scanner.Text()
		half := len(line) / 2 
		letters := make([]rune, half)

		var match rune
		for index, letter := range line {
			if index < half {
				letters[index] = letter
			} else {
				// Find duplicate character
				for _, p := range letters {
					if p == letter {
						match = letter
						break
					}
				}
			}
		}

		// Calculate priority
		// Add to total
		totalPriority += convertRune(match)
	}

	fmt.Println("Total priority:", totalPriority)
}

func convertRune(r rune) int32 {
	if r > 90 {
		return r - 96
	} else {
		return r - 38
	}
}