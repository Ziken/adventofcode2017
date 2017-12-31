package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

const INPUT_FILE = "input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func getInput() (input [][]string) {
	file, err := os.Open(INPUT_FILE)
	check(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		input = append(input, strings.Split(row, " "))
	}

	return
}
func answerPart1And2(instructions [][]string) {
	vars := make(map[string]int)
	var maxValueEver int
	for _, inst := range instructions {
		variable, action, amount, ifVar, condition, thisVal :=
		inst[0], inst[1], inst[2],inst[4], inst[5], inst[6]
		if checkCondition(vars, ifVar, condition, thisVal) {
			executeInstruction(vars, variable, action, amount)
			m := getMaxValue(vars)
			if m > maxValueEver {
				maxValueEver = m
			}
		}
	}
	fmt.Println("Answer part 1:", getMaxValue(vars))
	fmt.Println("Answer part 2:", maxValueEver)
}
func checkCondition(vars map[string]int, ifVar, condition, thisVal string) bool {
	val, errAtoi := strconv.Atoi(thisVal)
	check(errAtoi)
	switch condition {
		case ">":
			if vars[ifVar] > val {
				return true
			}
		case ">=":
			if vars[ifVar] >= val {
				return true
			}
		case "<":
			if vars[ifVar] < val {
				return true
			}
		case "<=":
			if vars[ifVar] <= val {
				return true
			}
		case "==":
			if vars[ifVar] == val {
				return true
			}
		case "!=":
			if vars[ifVar] != val {
				return true
			}
	}
	return false
}
func executeInstruction(vars map[string]int, variable, action, amount string) {
	val, errAtoi := strconv.Atoi(amount)
	check(errAtoi)
	switch action {
		case "inc":
			vars[variable] += val
		case "dec":
			vars[variable] -= val
	}
}
func getMaxValue(vars map[string]int) (max int){
	for _, value := range vars {
		if value > max {
			max = value
		}
	}
	return
}
func main() {
	in := getInput()
	answerPart1And2(in)
}