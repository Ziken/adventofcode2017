package main

import (
	"os"
	"bufio"
	"regexp"
	"fmt"
	"strconv"
)

const INPUT_FILE = "input.txt"

type Program struct {
	name string
	value int
	next []string
}

func check (e error) {
	if e != nil {
		panic(e)
	}
}
func getInput() map[string]Program {
	in := make(map[string]Program)
	file, errFile := os.Open(INPUT_FILE)
	check(errFile)
	scanner := bufio.NewScanner(file)
	check(scanner.Err())

	splitRegexp := regexp.MustCompile(`([a-z0-9]+)`)
	for scanner.Scan() {
		row := scanner.Text()
		parts := splitRegexp.FindAllString(row, -1)
		val, _ := strconv.Atoi(parts[1])
		in[parts[0]] = Program{
			name:parts[0],
			value:val,
			next: parts[2:],
			}
	}

	return in
}
func getBottomProgramName(programs map[string]Program) string {
	occursOfPrograms := make(map[string]int)
	for name, p := range programs {
		occursOfPrograms[name]++
		for _, pNext := range p.next {
			occursOfPrograms[pNext]++
		}
	}
	for p, value := range occursOfPrograms {
		if value == 1 {
			return p
		}
	}
	return "The program does not exist."
}
func answerPart2(programs map[string]Program, bottomProgramName string) {
	sumOfSubPrograms := make(map[string]int)
	getSumOfSubPrograms(programs, sumOfSubPrograms, bottomProgramName)
	fmt.Println("Answer part 2:", sumOfSubPrograms["ANSWER"])
}
func getSumOfSubPrograms (programs map[string]Program, sumOfPrograms map[string]int, subProgramName string) int {
	currentProg := programs[subProgramName]
	sum := currentProg.value
	if len(currentProg.next) == 0 {
		return sum
	}
	var arrOfSums [] int // sum of sub-programs of current program
	var progWithMaxSumName string
	for _, name := range currentProg.next {
		s := getSumOfSubPrograms(programs, sumOfPrograms, name)
		if s > sumOfPrograms[progWithMaxSumName] {
			progWithMaxSumName = name
		}
		arrOfSums = append(arrOfSums, s)
		sum += s
	}
	// Do not compute an value if this node is last or an answer is calculated
	if len(programs[progWithMaxSumName].next) != 0 && sumOfPrograms["ANSWER"]  <= 0 {
		for _, value := range arrOfSums {
			if sumOfPrograms[progWithMaxSumName] - value != 0 {
				sumOfPrograms["ANSWER"] = programs[progWithMaxSumName].value - (sumOfPrograms[progWithMaxSumName] - value)
			}
		}
	}
	sumOfPrograms[subProgramName] = sum
	return sum
}
func main() {
	in := getInput()
	bottomProgramName := getBottomProgramName(in)
	fmt.Println("Answer part 1:", bottomProgramName)
	answerPart2(in, bottomProgramName)
}