package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	//"strconv"
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

	hasReadInitialState := false
	var initialStateStrings = make([]string, 32) // Guestimate capacity

	lineIdx := 0
	scanner := bufio.NewScanner(file)
	var state []Stack
	for scanner.Scan() {
		// Process until we hit a newline, storing each string in an array
		lineText := scanner.Text()
		if !hasReadInitialState && len(lineText) > 0 {
			initialStateStrings[lineIdx] = lineText
			lineIdx++

			continue
		}

		// Once we hit a newline, process the array to generate our initial structure
		if !hasReadInitialState {
			hasReadInitialState = true

			state = processInitialState(initialStateStrings, lineIdx-1)
			continue
		}

		// Then, process each instruction line to modify the collection
		processInstruction(state, lineText)
	}

	// Read the final state of the collection
	for _, e := range state {
		fmt.Print(e.Pop())
	}
}

func processInitialState(inputStrings []string, lastIdx int) []Stack {
	// Last line tells us how many Stacks we need
	stackCountString := inputStrings[lastIdx]
	stackCountArray := strings.Split(stackCountString, "   ")
	stackCount := len(stackCountArray)

	output := make([]Stack, stackCount)

	for idx := lastIdx - 1; idx >= 0; idx-- {
		inputLine := inputStrings[idx]
		inputItems := make([]string, stackCount)

		for i := 0; i < len(inputLine); i += 4 {
			boundary := i + 4
			if boundary > len(inputLine) {
				boundary = len(inputLine)
			}

			inputItems[i/4] = strings.TrimSpace(inputLine[i:boundary])
		}

		stackIdx := 0
		for _, inp := range inputItems {

			if len(inp) == 0 {
				stackIdx++
				continue
			}

			output[stackIdx].Push(inp)
			stackIdx++
		}
	}

	return output
}

func processInstruction(state []Stack, instruction string) {
	instData := strings.Split(instruction, " ")
	quantity, err := strconv.Atoi(instData[1])
	if err != nil {
		handleError(err)
	}

	source, err := strconv.Atoi(instData[3])
	if err != nil {
		handleError(err)
	}

	target, err := strconv.Atoi(instData[5])
	if err != nil {
		handleError(err)
	}

	for i := 0; i < quantity; i++ {
		val := state[source - 1].Pop()
		state[target - 1].Push(val)
	}
}