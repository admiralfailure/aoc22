package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	Run()
}

func Run() {
	file, err := os.Open("input.txt")
	//defer file.Close()

	if err != nil {
		panic(err)
	}

	runningCount := 0
	var totals [3]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			if runningCount > totals[0] {
				totals[2] = totals[1]
				totals[1] = totals[0]
				totals[0] = runningCount
			} else if runningCount > totals[1] {
				totals[2] = totals[1]
				totals[1] = runningCount
			} else if runningCount > totals[2] {
				totals[2] = runningCount
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

	fmt.Println("1st elf: ", totals[0])
	fmt.Println("2nd elf: ", totals[1])
	fmt.Println("3rd elf: ", totals[2])
	fmt.Println()
	fmt.Println("Total: ", totals[0]+totals[1]+totals[2])
}
