package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

const INPUT_FILE = "input.txt"
// size of graph
const SIZE = 50 + 1

func check (e error) {
	if e != nil {
		panic(e)
	}
}
func getInput() (adjacencyMatrix [][]byte) {
	for i := 0; i < SIZE; i++ {
		adjacencyMatrix = append(adjacencyMatrix, make([]byte, SIZE))
	}
	file, errFile := os.Open(INPUT_FILE)
	check(errFile)
	scanner := bufio.NewScanner(file)
	check(scanner.Err())
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "/")
		i,_ := strconv.Atoi(row[0])
		j,_ := strconv.Atoi(row[1])

		adjacencyMatrix[i][j] = 1
		adjacencyMatrix[j][i] = 1
	}
	return
}
func answerPart1And2(in [][]byte) {
	var bridges [][]int
	adjacencyMatrix := copyAdjacencyMatrix(in)
	goThroughGraph(adjacencyMatrix, 0, &bridges,0,0)
	var maxStrong, longest, longestAndStrongest int
	for _, bridge := range bridges {
		strong, size := bridge[0], bridge[1]
		if strong > maxStrong {
			maxStrong = strong
		}
		if size > longest {
			longest = size
			longestAndStrongest = longest
		} else if size == longest && strong > longestAndStrongest {
			longestAndStrongest = strong
		}
	}
	fmt.Println("Answer part 1:", maxStrong)
	fmt.Println("Answer part 2:", longestAndStrongest)
}
func copyAdjacencyMatrix(adjacencyMatrix [][]byte) [][]byte {
	c := make([][] byte, SIZE)
	for i := 0; i < SIZE; i++ {
		c[i] = make([]byte, SIZE)
		copy(c[i], adjacencyMatrix[i])
	}
	return c
}
func goThroughGraph(adjacencyMatrix [][]byte, node int, bridges *[][]int, currentSum, lenBridge int) {
	for i := 0; i < SIZE; i++ {
		if adjacencyMatrix[i][node] == 1 {
			c := copyAdjacencyMatrix(adjacencyMatrix)
			c[i][node] = 0
			c[node][i] = 0
			*bridges = append(*bridges, []int{currentSum+i+node, lenBridge+1})
			goThroughGraph(c, i, bridges, currentSum+i+node, lenBridge+1)
		}
	}
}
func main() {
	in := getInput()
	answerPart1And2(in)
}