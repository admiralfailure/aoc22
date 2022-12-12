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

	var routeMap []rune
	mapWidth := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()

		// Read input into array
		// Get width
		if mapWidth == 0 {
			mapWidth = len(input)
			routeMap = make([]rune, 0, 128*mapWidth)
		}

		// Split input
		for _, elm := range input {
			// col is X, row is Y
			// idx := (x * mapWidth) + y
			height := elm
			if err != nil {
				handleError(err)
			}

			routeMap = append(routeMap, height) // Append is as fast as index assign if initial
			// capacity is great enough
		}
	}

	//fmt.Println(routeMap)

	// Process map and create structs
	nodeMap := make([]node, len(routeMap))
	for idx, elm := range routeMap {
		nodeMap[idx] = node{name: elm}
	}

	// Process neighbours
	for idx, _ := range nodeMap {
		elm := &nodeMap[idx]
		// Get valid routes out
		validRoutes := getValidRoutes(routeMap, idx, mapWidth)

		elm.neighbours = make([]int, len(validRoutes))
		for idx2, elm2 := range validRoutes {
			//fmt.Println("adding neighbour to", string(elm.name), "-", string(nodeMap[elm2].name))
			elm.neighbours[idx2] = elm2
		}
	}

	//fmt.Println(nodeMap)

	source := 0
	target := 0
	for i := range nodeMap {
		if nodeMap[i].name == 'S' {
			source = i
		}
		if nodeMap[i].name == 'E' {
			target = i
		}
	}

	shortestPath := dijkstra(nodeMap, source, target)
	fmt.Println("Shortest path:", shortestPath)
}

func getValidRoutes(routeMap []rune, idx int, mapWidth int) []int {
	// Get neighbour indices
	x := idx % (mapWidth)
	y := idx / (mapWidth)
	mapHeight := len(routeMap) / mapWidth

	//fmt.Println(idx, x, y, string(routeMap[idx]))

	// Get neighbour indices
	var idxUp, idxDown, idxLeft, idxRight int
	if x%mapWidth > 0 {
		idxLeft = idx - 1
	} else {
		idxLeft = -1
	}
	if x%mapWidth < mapWidth-1 {
		idxRight = idx + 1
	} else {
		idxRight = -1
	}
	if y > 0 {
		idxUp = idx - mapWidth
	} else {
		idxUp = -1
	}
	if y < mapHeight-1 {
		idxDown = idx + mapWidth
	} else {
		idxDown = -1
	}

	//fmt.Println(idxLeft, idxRight, idxUp, idxDown)

	// Work out if the routes are valid
	validRoutes := make([]int, 0, 4)
	if idxLeft >= 0 && isRouteValid(routeMap[idx], routeMap[idxLeft]) {
		validRoutes = append(validRoutes, idxLeft)
	}
	if idxRight >= 0 && isRouteValid(routeMap[idx], routeMap[idxRight]) {
		validRoutes = append(validRoutes, idxRight)
	}
	if idxUp >= 0 && isRouteValid(routeMap[idx], routeMap[idxUp]) {
		validRoutes = append(validRoutes, idxUp)
	}
	if idxDown >= 0 && isRouteValid(routeMap[idx], routeMap[idxDown]) {
		validRoutes = append(validRoutes, idxDown)
	}

	return validRoutes
}

func isRouteValid(source rune, target rune) bool {
	if source == 'S' || target == 'S' || source == 'E' || target == 'E' {
		fmt.Println("Source", string(source), "Target", string(target))

		if target == 'S' || (source == 'S' && (target == 'a' || target == 'b')) {
			fmt.Println("case 1")
			return true
		}
	
		if target == 'E' && (source == 'z' || source == 'y') {
			fmt.Println("case 2")
			return true
		}

		fmt.Println("no match")
		return false
	}

	if target-source == 1 {
		return true
	}

	if source >= target {
		return true
	}

	return false
}

func dijkstra(graph []node, source int, target int) int {
	fmt.Println("Dijkstra:", source, string(graph[source].name), target, string(graph[target].name))

	dist := make([]int, len(graph))
	visited := make([]bool, len(graph))

	for idx := range graph {
		dist[idx] = 99999
		visited[idx] = false
	}

	dist[source] = 0
	//visited[source] = true
	nextNode := 0
	count := 1

	//fmt.Println("Processing", dist, visited)
	for count < len(graph) {
		minDistance := 99999
		for i := 0; i < len(graph); i++ {
			// Get closest node
			if dist[i] < minDistance && !visited[i] {
				minDistance = dist[i]
				nextNode = i
			}
		}

		visited[nextNode] = true
		
		for _, e := range graph[nextNode].neighbours {
			if !visited[e]{
				if dist[e] > minDistance+1 {				
					fmt.Println("Setting distance from", string(graph[nextNode].name), "to", string(graph[e].name), "to", minDistance+1)
					dist[e] = minDistance + 1
				}
			}
		}

		count++
	}

	return dist[target]
}
