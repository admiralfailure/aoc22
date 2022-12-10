package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

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

	head := point{x: 0, y: 0}
	tail := point{x: 0, y: 0}
	visited := make(map[string]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := scanner.Text()

		// Move head
		instructionParts := strings.Split(instruction, " ")
		direction := instructionParts[0]
		amount, err := strconv.Atoi(instructionParts[1])
		if err != nil {
			handleError(err)
		}

		for i := amount; i > 0; i-- {
			moveHead(&head, direction)

			fmt.Println("Head position:", head.x, ",", head.y)

			moveTail(head, &tail)
			fmt.Println("Tail position:", tail.x, ",", tail.y)

			key := strconv.Itoa(tail.x) + "|" + strconv.Itoa(tail.y)
			if _, ok := visited[key]; !ok {
				visited[key] = true
			}
		}
	}

	fmt.Println("Total visited:", len(visited))
}

func moveHead(head *point, direction string) {
	switch direction {
	case "U":
		head.y++
	case "D":
		head.y--
	case "R":
		head.x++
	case "L":
		head.x--
	}
}

func moveTail(head point, tail *point) {
	// Same place? No move
	if head.x == tail.x && head.y == tail.y {
		fmt.Println("NO MOVE: On top")
		return
	}

	// Within one in any direction? No move
	if math.Abs(float64(head.x-tail.x)) <= 1 && math.Abs(float64(head.y-tail.y)) <= 1 {
		fmt.Println("NO MOVE: Within one")
		return
	}

	// CARDINAL DIRECTIONS
	// UP
	if head.y-tail.y > 1 && head.x == tail.x {
		fmt.Println("UP")
		tail.y++
		return
	}
	// DOWN
	if tail.y-head.y > 1 && head.x == tail.x {
		fmt.Println("DOWN")
		tail.y--
		return
	}
	// RIGHT
	if head.x-tail.x > 1 && head.y == tail.y {
		fmt.Println("RIGHT")
		tail.x++
		return
	}
	// LEFT
	if tail.x-head.x > 1 && head.y == tail.y {
		fmt.Println("LEFT")
		tail.x--
		return
	}

	// DIAGONALS
	// UP RIGHT
	if head.x > tail.x && head.y > tail.y {
		fmt.Println("UP RIGHT")
		tail.y++
		tail.x++
		return
	}
	// DOWN RIGHT
	if head.x > tail.x && head.y < tail.y {
		fmt.Println("DOWN RIGHT")
		tail.y--
		tail.x++
		return
	}
	// DOWN LEFT
	if head.x < tail.x && head.y < tail.y {
		fmt.Println("DOWN LEFT")
		tail.y--
		tail.x--
		return
	}
	// UP LEFT
	if head.x < tail.x && head.y > tail.y {
		fmt.Println("UP LEFT")
		tail.y++
		tail.x--
		return
	}
}
