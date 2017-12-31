package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

const INPUT_FILE = "input.txt"

func check (e error) {
	if e != nil {
		panic(e)
	}
}
func getInput() (input [][]int) {
	indexRow := 0
	file, errFile := os.Open(INPUT_FILE)
	check(errFile)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	check(scanner.Err())
	for scanner.Scan() {
		input = append(input, make([]int, 0))

		row := scanner.Text()
		splittedRow := strings.Split(row, " <-> ")

		programs := strings.Split(splittedRow[1], ", ")
		for _, p := range programs {
			n, errAtoi := strconv.Atoi(p)
			check(errAtoi)
			input[indexRow] = append(input[indexRow], n)
		}
		indexRow++
	}
	return input
}

func answerPart1 (input [][]int) {
	amountOfPrograms := len(input)
	isVisited := make([]bool, amountOfPrograms)

	// check only connections of first program
	dfs(0,isVisited,input)

	sumConnectedTo0 := 0
	for _,isConnected := range isVisited {
		if isConnected {
			sumConnectedTo0++
		}
	}
	fmt.Println("Answer part 1:", sumConnectedTo0)
}
func dfs(s int, isVisited []bool, input [][]int) {
	isVisited[s] = true
	for a := 0; a < len(input[s]); a++ {
		if !isVisited[input[s][a]] {
			dfs(input[s][a], isVisited, input)
		}
	}
}
func answerPart2(input [][]int) {
	amountOfPrograms := len(input)
	isVisited := make([]bool, amountOfPrograms)
	amountOfGroups := 0
	for i,vis := range isVisited {
		if !vis {
			amountOfGroups++
			dfs(i, isVisited, input)
		}
	}
	fmt.Println("Answer part 2:", amountOfGroups)
}
func main() {
	in := getInput()
	answerPart1(in)
	answerPart2(in)
}
