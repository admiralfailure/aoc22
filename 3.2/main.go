package main

import (
	"bufio"
	"fmt"
	"os"
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

	totalPriority := int32(0)
	scanner := bufio.NewScanner(file)

	var lines [3]string
	lineIdx := 0
	for scanner.Scan() {
		// Push to the collection
		lines[lineIdx] = scanner.Text()

		if lineIdx == 2 {
			// Process
			match := process(lines)

			// Calculate priority
			// Add to total
			totalPriority += convertRune(match)

			// Clear collection
			for i, _ := range lines {
				lines[i] = ""
			}

			// Clear index
			lineIdx = 0
		} else {
			lineIdx++
		}
	}

	fmt.Println("Total priority:", totalPriority)
}

func process(lines [3]string) rune {
	fmt.Println("Processing...")
	fmt.Println(lines[0])
	fmt.Println(lines[1])
	fmt.Println(lines[2])

	match := getMatchingRunes(lines[0], lines[1], lines[2])

	fmt.Println("Match:", string(match))
	fmt.Println()

	return match
}

func getMatchingRunes(s1 string, s2 string, s3 string) rune {
	for _, r1 := range s1 {
		// Is it in s2?
		hasMatched := false
		for _, r2 := range s2 {
			if r1 == r2 {
				hasMatched = true
				continue
			}
		}

		if !hasMatched {
			continue
		}

		// Is it in s3?
		for _, r3 := range s3 {
			if r1 == r3 {
				return r1
			}
		}
	}

	return rune(0)
}

func convertRune(r rune) int32 {
	if r > 90 {
		return r - 96
	} else {
		return r - 38
	}
}
