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

	xReg := 1
	cycle := 0
	totalSignalStrength := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := scanner.Text()
		instructionParts := strings.Split(instruction, " ")
		command := instructionParts[0]

		switch command {
		case "noop":
			totalSignalStrength += runCycle(&cycle, xReg)
			continue
		case "addx":
			totalSignalStrength += runCycle(&cycle, xReg)
			totalSignalStrength += runCycle(&cycle, xReg)

			value, err := strconv.Atoi(instructionParts[1])
			if err != nil {
				handleError(err)
			}

			xReg += value
		}
	}

	fmt.Println("Total Signal Strength", totalSignalStrength)
}

func runCycle(cycle *int, xReg int) int {
	*cycle++

	if (*cycle-20)%40 == 0 {
		signalStrength := *cycle * xReg
		fmt.Println("Cycle", *cycle, "Value", xReg, "Signal Strength", signalStrength)
		return signalStrength
	}

	return 0
}
