package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"strconv"
)

const INPUT_FILE = "input.txt"

func check (e error) {
	if e != nil {
		panic(e)
	}
}
func getInput() map[int][]string {
	in := make(map[int][]string)
	file, errFile := os.Open(INPUT_FILE)
	check(errFile)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	check(scanner.Err())
	indexRow := 0
	for scanner.Scan() {
		row := scanner.Text()
		in[indexRow] = strings.Split(row, " ")
		indexRow++
	}
	return in
}
func divideInstruction(instruction []string, register map[string]int) (command, l string, value int) {
	command, l, value = instruction[0], instruction[1], 0
	if len(instruction) > 2 {
		// if Atoi returns error (errAtoi) it means that third value is not an integer
		// it is the key for value stored in "register"
		if val, errAtoi := strconv.Atoi(instruction[2]); errAtoi != nil {
			value = register[ instruction[2] ]
		} else {
			value = val
		}
	}
	return
}
func answerPart1 (instructions map[int][]string) {
	register := make(map[string]int)
	mulOperationAmount := 0
	for i := 0; i < len(instructions); i++ {

		command, l, value := divideInstruction(instructions[i], register)

		switch command {
			case "set":
				register[l] = value
			case "sub":
				register[l] -= value
			case "mul": // multiply
				register[l] *= value
				mulOperationAmount++
			case "jnz": // jumps
				if v, err := strconv.Atoi(l); err == nil {
					if v != 0 {
						i += value - 1
					}
				} else if register[l] != 0 {
					i += value - 1
				}

		} // switch
	} // for

	fmt.Println("Answer part 1:", mulOperationAmount)
}
func getBandC (instructions map[int][]string) (int,int) {
	register := make(map[string]int)
	register["a"] = 1
	for i := 0; i < 8; i++ {

		command, l, value := divideInstruction(instructions[i], register)

		switch command {
			case "set":
				register[l] = value
			case "sub":
				register[l] -= value
			case "mul": // multiply
				register[l] *= value
			case "jnz": // jumps
				if v, err := strconv.Atoi(l); err == nil {
					if v != 0 {
						i += value - 1
					}
				} else if register[l] != 0 {
					i += value - 1
				}
		} // switch
	} // for
	return register["b"],register["c"]
}
func answerPart2 (instructions map[int][]string) {
	b,c := getBandC(instructions)
	sum := 0
	for ; b <= c; b += 17 {
		for i := 2; i < b; i++ {
			if b % i == 0 {
				sum++
				break
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