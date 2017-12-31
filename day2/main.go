package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
	//"math"
	"sort"
)

const INPUT_FILE = "input.txt"

func check (e error) {
	if e != nil {
		panic(e)
	}
}
func getInput() [][]int {
	input := make([][]int,0)
	indexRow := 0
	file, errFile := os.Open(INPUT_FILE)
	check(errFile)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	check(scanner.Err())
	for scanner.Scan() {
		var in []int
		row := scanner.Text()
		splittedRow := strings.Split(row, "\t")

		for _, p := range splittedRow {
			n, errAtoi := strconv.Atoi(p)
			check(errAtoi)
			in = append(in, n)
		}
		input = append(input, in)
		indexRow++
	}
	return input
}
func answerPart1(numbers [][]int) {
	var checksum int

	for _,row := range numbers {
		sort.Ints(row)
		checksum += row[len(row)-1] - row[0]
	}
	fmt.Println("Answer part 1:", checksum)
}
func answerPart2(numbers [][]int) {
	var sum int

	for _,row := range numbers {

		for i, n := range row {
			for j := i+1; j < len(row); j++ {
				if n % row[j] == 0 {
					sum += n/row[j]
				} else if row[j] % n == 0 {
					sum += row[j]/n
				}
			}
		}

	}
	fmt.Println("Answer part 2:", sum)
}
func main() {
	in := getInput()
	answerPart1(in)
	answerPart2(in)
}
