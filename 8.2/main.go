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

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func Run() {
	file, err := os.Open("input.txt")
	handleError(err)

	defer file.Close()

	var treeMap []int
	mapWidth := 0
	//mapHeight := 0
	row := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()

		// Read input into array
		// Get width
		if mapWidth == 0 {
			mapWidth = len(input)
			treeMap = make([]int, 0, 128*mapWidth)
		}

		// Split input
		for _, elm := range input {
			// col is X, row is Y
			// idx := (x * mapWidth) + y
			height, err := strconv.Atoi(string(elm))
			if err != nil {
				handleError(err)
			}

			treeMap = append(treeMap, height) // Append is as fast as index assign if initial
			// capacity is great enough
		}

		// Increment row
		row++
	}

	// Work on array
	//fmt.Println(treeMap)

	maxScore := 0
	for idx, _ := range treeMap {

		scenicScore := getTreeScenicScore(treeMap, idx, mapWidth-1, row-1)
		if scenicScore > maxScore {
			maxScore = scenicScore
		}
	}

	fmt.Println("Max score:", maxScore)
}

func getIndex(x int, y int, mapWidth int) int {
	return (y * (mapWidth + 1)) + x
}

func getTreeScenicScore(treeMap []int, idx int, mapWidth int, mapHeight int) int {
	x := idx % (mapWidth + 1)
	y := idx / (mapWidth + 1)
	val := treeMap[idx]

	fmt.Print("Checking", val, "...")

	// Edges have a 0 view so always return 0
	if x == 0 || y == 0 || x == mapWidth || y == mapHeight {
		fmt.Print("Edge\n")
		return 0
	}

	// Otherwise check each cardinal direction
	// Up
	upScore := 0
	for dY := 1; y-dY >= 0; dY++ {
		curY := y - dY
		if treeMap[getIndex(x, curY, mapWidth)] >= val {
			upScore++
			break
		}

		//fmt.Print("Up\n")
		upScore++
	}
	fmt.Print("Up:", upScore)

	// Down
	downScore := 0
	for dY := 1; y+dY <= mapWidth; dY++ {
		curY := y + dY
		if treeMap[getIndex(x, curY, mapWidth)] >= val {
			downScore++
			break
		}

		//fmt.Print("Down\n")
		downScore++
	}
	fmt.Print("Down:", downScore)

	// Left
	leftScore := 0
	for dX := 1; x-dX >= 0; dX++ {
		curX := x - dX

		if treeMap[getIndex(curX, y, mapWidth)] >= val {
			leftScore++
			break
		}

		//fmt.Print("Left\n")
		leftScore++
	}
	fmt.Print("Left:", leftScore)

	// Right
	rightScore := 0
	for dX := 1; x+dX <= mapHeight; dX++ {
		curX := x + dX

		if treeMap[getIndex(curX, y, mapWidth)] >= val {
			rightScore++
			break
		}

		//fmt.Print("Right\n")
		rightScore++
	}
	fmt.Print("Right:", rightScore)

	// Reached the end
	totalScore := upScore * downScore * leftScore * rightScore
	fmt.Print("Total:", totalScore, "\n")
	return totalScore
}
