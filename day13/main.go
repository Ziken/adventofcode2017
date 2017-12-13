package main

import (
	"strings"
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
func getInput() map[int]int {
	input := make(map[int]int)

	file, errFile := os.Open(INPUT_FILE)
	check(errFile)
	scanner := bufio.NewScanner(file)
	check(scanner.Err())
	for scanner.Scan() {
		row := scanner.Text()
		splitedRow := strings.Split(row, ": ")
		layer, errLayer := strconv.Atoi(splitedRow[0])
		check(errLayer)
		rangeLayer, errRangeLayer := strconv.Atoi(splitedRow[1])
		check(errRangeLayer)
		input[layer] = rangeLayer
	}
	return input
}
func answerPart1 (in map[int]int) {
	answer := 0
	for layer, r := range in {
		// definition of Arithmetic progression :)
		// range                | 2 | 3 | 4 | 5 | 6
		// steps for one cycle  | 2 | 4 | 6 | 8 | 10
		if  layer % (2*r-2) == 0 {
			answer += layer*r
		}
	}
	fmt.Println("Part 1:", answer)
}
func answerPart2 (in map[int]int) {
	delay := 0
	for {
		isCaught := false
		for layer, r := range in {
			if  (layer+delay) % (2*r-2) == 0 {
				isCaught = true
				break
			}
		}
		if !isCaught {
			break
		}
		delay++
	}
	fmt.Println("Part 2:", delay)

}
func main () {
	in := getInput()
	answerPart1(in)
	answerPart2(in)

}
