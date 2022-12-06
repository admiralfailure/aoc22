package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	//"strconv"
	//"strings"
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := []int32(scanner.Text())
		
		for i:= 3; i < len(input); i++ {
			buf := make([]int32, 4)
			copy(buf, input[i-3 : i+1])

			//fmt.Println(string(buf))
			sort.Slice(buf, func(i int, j int) bool { return buf[i] < buf[j]; })
			//fmt.Println(string(buf))


			if isDistinct(buf) {
				fmt.Println("First distinct index is", i + 1)
				return
			}
			
		}
	}
}

func isDistinct(buf []rune) bool {
	for i := 0; i < len(buf) - 1; i++ {
		if buf[i] == buf[i + 1] {
			return false
		}
	}

	return true
}