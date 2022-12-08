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

	var tree TreeNode = &DirNode{Name: "/"}
	root := tree

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()

		// Process line
		// Can either be instruction or output
		if strings.HasPrefix(input, "$") {
			// If instruction, update our current location accordingly
			instructionParts := strings.Split(input, " ")

			switch instructionParts[1] {
			case "cd":
				switch instructionParts[2] {
				case "/":
					tree = root
					break
				case "..":
					tree = tree.GetParent()
					break
				default:
					for _, elm := range tree.Children() {
						if elm.GetName() == instructionParts[2] {
							tree = elm
							break
						}
					}

					break
				}
				break
			case "ls":
				break
			}
		} else {
			// If output, update data accordingly
			dataParts := strings.Split(input, " ")
			if dataParts[0] == "dir" {
				// Directory
				newNode := DirNode{Name: dataParts[1]}
				newNode.Parent = tree

				//fmt.Println("Adding child", newNode, "to tree", tree)
				tree.AddChild(&newNode)
			} else {
				// File
				size, _ := strconv.Atoi(dataParts[0])
				newNode := FileNode{Name: dataParts[1], FileSize: uint64(size)}
				newNode.Parent = tree

				tree.AddChild(&newNode)
			}
		}

	}

	var fsSize uint64 = 70000000
	var requiredUnused uint64 = 30000000
	availableSpace := fsSize - root.Size()
	requiredExtra := requiredUnused - availableSpace

	fmt.Println("Smallest deletion candidate:", getSmallestDeletionCandidate(root, requiredExtra))
}

func getSmallestDeletionCandidate(node TreeNode, requiredSpace uint64) uint64 {
	var smallest uint64 = 0

	if len(node.Children()) > 0 && node.Size() >= requiredSpace {
		if smallest == 0 || node.Size() < smallest {
			smallest = node.Size()
		}
	}

	for _, elm := range node.Children() {
		elmTotal := getSmallestDeletionCandidate(elm, requiredSpace)
		if smallest == 0 || (elmTotal > 0 && elmTotal < smallest) {
			smallest = elmTotal
		}
	}

	return smallest
}
