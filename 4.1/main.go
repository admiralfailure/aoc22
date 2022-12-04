package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	inclusivePairs := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Process lines into pair strings
		line := scanner.Text()
		p1, p2 := processLine(line)

		// Process pair strings into numbers
		n1, n2 := processPair(p1)
		n3, n4 := processPair(p2)

		// Check inclusion
		r1 := n2 - n1
		r2 := n4 - n3

		if r1 > r2 {
			if n3 >= n1 && n4 <= n2 {
				inclusivePairs++
			}
		} else if r2 > r1 {
			if n1 >= n3 && n2 <= n4 {
				inclusivePairs++
			}
		} else {
			if n1 == n3 {
				inclusivePairs++
			}
		}

		//fmt.Println()
	}

	fmt.Println("Total inclusive pairs:", inclusivePairs)
}

func processLine(line string) (string, string) {
	// Split on ,
	data := strings.Split(line, ",")

	// Return
	return data[0], data[1]
}

func processPair(pair string) (int, int) {
	// Split on -
	data := strings.Split(pair, "-")

	// Convert
	n1, e1 := strconv.Atoi(data[0])
	handleError(e1)
	n2, e2 := strconv.Atoi(data[1])
	handleError(e2)

	// Return
	return n1, n2
}
