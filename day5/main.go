package main

import (
	"os"
	"bufio"
	"strconv"
	"fmt"
)

const INPUT_FILE = "input.txt"

func check (e error) {
	if e != nil {
		panic(e)
	}
}
func getInput() (in []int) {
	file, errFile := os.Open(INPUT_FILE)
	check(errFile)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	check(scanner.Err())
	for scanner.Scan() {
		row := scanner.Text()
		n, errAtoi := strconv.Atoi(row)
		check(errAtoi)
		in = append(in, n)
	}
	return
}
func answerPart1(in []int) {
	maze := make([]int, len(in))
	copy(maze, in)
	steps := 0
	currentPos := 0
	for currentPos < len(maze) {
		currentPos, maze[currentPos] = currentPos+maze[currentPos], maze[currentPos]+1
		steps++
	}
	fmt.Println("Answer part 1:", steps)
}
func answerPart2(in []int) {
	maze := make([]int, len(in))
	copy(maze, in)
	steps := 0
	currentPos := 0
	for currentPos < len(maze) {
		if maze[currentPos] >= 3 {
			currentPos, maze[currentPos] = currentPos+maze[currentPos], maze[currentPos] - 1
		} else {
			currentPos, maze[currentPos] = currentPos+maze[currentPos], maze[currentPos] + 1
		}
		steps++
	}
	fmt.Println("Answer part 2:", steps)
}
func main() {
	in := getInput()
	answerPart1(in)
	answerPart2(in)
}