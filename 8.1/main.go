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

	visibleTrees := 0
	for idx, _ := range treeMap {
		if isTreeVisible(treeMap, idx, mapWidth-1, row-1) {
			visibleTrees++
		}
	}

	fmt.Println("Visible trees:", visibleTrees)
}

func getIndex(x int, y int, mapWidth int) int {
	return (y * (mapWidth + 1)) + x
}

func isTreeVisible(treeMap []int, idx int, mapWidth int, mapHeight int) bool {
	x := idx % (mapWidth + 1)
	y := idx / (mapWidth + 1)
	val := treeMap[idx]

	//fmt.Print("Checking", val, "...")

	// Edges are always visible
	if x == 0 || y == 0 || x == mapWidth || y == mapHeight {
		//fmt.Print("Edge\n")
		return true
	}

	// Otherwise check each cardinal direction
	// Up
	for dY := 1; y-dY >= 0; dY++ {
		curY := y - dY
		if treeMap[getIndex(x, curY, mapWidth)] >= val {
			break
		}

		// Reached the edge with no higher trees
		if curY == 0 {
			//fmt.Print("Up\n")
			return true
		}
	}

	// Down
	for dY := 1; y+dY <= mapWidth; dY++ {
		curY := y + dY
		if treeMap[getIndex(x, curY, mapWidth)] >= val {
			break
		}

		// Reached the edge with no higher trees
		if curY == mapWidth {
			//fmt.Print("Down\n")
			return true
		}
	}

	// Left
	for dX := 1; x-dX >= 0; dX++ {
		curX := x - dX

		if treeMap[getIndex(curX, y, mapWidth)] >= val {
			break
		}

		// Reached the edge with no higher trees
		if curX == 0 {
			//fmt.Print("Left\n")
			return true
		}
	}

	// Right
	for dX := 1; x+dX <= mapHeight; dX++ {
		curX := x + dX

		if treeMap[getIndex(curX, y, mapWidth)] >= val {
			break
		}

		// Reached the edge with no higher trees
		if curX == mapHeight {
			//fmt.Print("Right\n")
			return true
		}
	}

	// Reached the end
	//fmt.Print("Hidden\n")
	return false
}
