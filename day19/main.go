package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
)

const INPUT_FILE = "input.txt"

func check (e error) {
	if e != nil {
		panic(e)
	}
}
func getInput() (in [][]string) {

	file, errFile := os.Open(INPUT_FILE)
	check(errFile)
	scanner := bufio.NewScanner(file)
	check(scanner.Err())
	indexRow := 0
	for scanner.Scan() {
		row := scanner.Text()
		in = append(in, strings.Split(row, ""))
		indexRow++
	}
	return in
}
func answerPart1 (path [][]string) {
	var row, column int
	var stepVert, stepHoriz int
	var stepsAmount int
	var letters string
	column = findStart(path[0])
	stepVert = 1 // go down

	for {
		p := path[row][column]

		if p == "+" {
			stepVert, stepHoriz = setDirection(row,column,stepVert,stepHoriz,path)
		} else if isLetter(p) {
			letters += p
		} else if p != "|" &&  p != "-" {
			break
		}
		if ifCanMove(row,column,stepVert,stepHoriz,path) {
			stepsAmount++
			row += stepVert
			column += stepHoriz
		} else {
			stepsAmount++ // if last sign is letter
			break
		}
	}
	fmt.Println("Answer part 1:", letters)
	fmt.Println("Answer part 2:", stepsAmount)
}
func findStart(path []string) int {
	for i,v := range path {
		if v == "|" {
			return i
		}
	}
	return -1
}
func ifCanGo(direction string, row, column int, path [][]string) bool {
	switch direction {
	case "up":
		if row > 0 && (isLetter(path[row-1][column]) || path[row-1][column] == "|") {
			return true
		}
	case "down":
		if row < (len(path) - 1) && (isLetter(path[row+1][column]) || path[row+1][column] == "|"){
			return true
		}
	case "left":
		if column > 0 && (isLetter(path[row][column-1]) || path[row][column-1] == "-") {
			return true
		}
	case "right":
		if column < (len(path[row]) - 1) && (isLetter(path[row][column+1]) || path[row][column+1] == "-") {
			return true
		}
	}
	return false

}
func setDirection (row, column, stepVert, stepHoriz int, path [][]string) (int, int) {
	if ifCanGo("left", row, column, path) && stepHoriz == 0 {
		stepVert = 0
		stepHoriz = -1
	} else if ifCanGo("right", row, column, path) && stepHoriz == 0 {
		stepHoriz = 1
		stepVert = 0
	} else if ifCanGo("up", row, column, path) && stepVert == 0 {
		stepHoriz = 0
		stepVert = -1
	} else if ifCanGo("down", row, column, path) && stepVert == 0 {
		stepHoriz = 0
		stepVert = 1
	} else {
		stepHoriz = 2
		stepVert = 2
	}
	return stepVert, stepHoriz
}
func ifCanMove(row, column, stepVert, stepHoriz int, path [][]string) bool {
	if stepVert == -1 && row > 0 { // up
		return true
	} else if stepVert == 1 && row < (len(path) - 1) { // down
		return true
	} else if stepHoriz == 1 && column < (len(path[row]) - 1) { // right
		return true
	} else if stepHoriz == -1 && column > 0{ // left
		return true
	}
	return false
}
func isLetter(sign string) bool {
	if sign >= "A" && sign <= "Z"{
		return true
	}
	return false
}
func main() {
	in := getInput()
	//for _,v := range in {
	//	fmt.Println(v)
	//}
	answerPart1(in)
}
