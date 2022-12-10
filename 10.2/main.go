package main

import (
	"bufio"
	"fmt"
	"math"
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
	outputArray := make([]string, 40*6)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := scanner.Text()
		instructionParts := strings.Split(instruction, " ")
		command := instructionParts[0]

		switch command {
		case "noop":
			runCycle(&cycle, xReg, outputArray)
			continue
		case "addx":
			runCycle(&cycle, xReg, outputArray)
			runCycle(&cycle, xReg, outputArray)

			value, err := strconv.Atoi(instructionParts[1])
			if err != nil {
				handleError(err)
			}

			xReg += value
		}
	}

	//fmt.Println(outputArray)
	for y := 0; y < 6; y++ {
		for x := 0; x < 40; x++ {
			fmt.Print(outputArray[40*y+x])
		}
		fmt.Print("\n")
	}
}

func runCycle(cycle *int, xReg int, output []string) {
	*cycle++

	fmt.Println("Cycle:", *cycle)
	fmt.Println("Position:", *cycle%40-1)
	fmt.Println("Sprite position:", xReg-1, xReg+1)
	isSpriteVisible := math.Abs(float64(xReg-(*cycle%40-1))) <= 1
	//fmt.Println("IsVisible:", isSpriteVisible)
	fmt.Println()

	if isSpriteVisible {
		//fmt.Println("Setting # at", *cycle-1)
		output[*cycle-1] = "#"
	} else {
		//fmt.Println("Setting . at", *cycle-1)
		output[*cycle-1] = "."
	}

	//fmt.Println("End of cycle", *cycle, "xReg", xReg)
}
