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
func answerPart1 (instructions map[int][]string) {
	register := make(map[string][]int)
	recentlyPlayedSound := 0
	var endLoop bool
	for i:=0; i < len(instructions); i++ {

		command, l, value := instructions[i][0], instructions[i][1], 0
		if len(instructions[i]) > 2 {
			// if Atoi returns error (errAtoi) it means that third value is not an integer
			// it is the key for value stored in "register"
			if val, errAtoi := strconv.Atoi(instructions[i][2]); errAtoi != nil {
				value = register[ instructions[i][2] ][0]
			} else {
				value = val
			}
		}


		if len(register[l]) == 0 {
			// first index for current value, second one for recovery
			register[l] = append(register[l], 0,0)
		}

		switch command {
			case "set":
				register[l][0], register[l][1] = value, register[l][0]
			case "add":
					register[l][1] = register[l][0]
					register[l][0] += value
			case "mul": // multiply
					register[l][1] = register[l][0]
					register[l][0] *= value
			case "mod": // modulo
					register[l][1] = register[l][0]
					register[l][0] %= value
			case "rcv": // recover
				if register[l][0] != 0 {
					register[l][0] = register[l][1]
					endLoop = true
				}
			case "jgz": // jumps
				if register[l][0] > 0 {
					i += value-1
				}
			case "snd": // play a sound
				recentlyPlayedSound = register[l][0]
		} // switch

		if endLoop {
			break
		}

	} // for
	fmt.Println("Answer part 1:", recentlyPlayedSound)
}
type Program struct {
	register      map[string]int
	currentIndex  int
	sentValues    []int
	isWaiting     bool
	sentValAmount int
}
func (p * Program) initProgram(id int) {
	p.register["p"] = id
}
func (p * Program) execute(receivedValues * []int, instructions map[int][]string) {
	if p.currentIndex >= len(instructions) {
		p.isWaiting = true
		return
	}
	command, l := instructions[p.currentIndex][0], instructions[p.currentIndex][1]
	var value int
	if len(instructions[p.currentIndex]) > 2 {
		// if Atoi returns an error (errAtoi) it means that third value is not an integer
		// it is the key for value stored in "register"
		if val, errAtoi := strconv.Atoi(instructions[p.currentIndex][2]); errAtoi != nil {
			value = p.register[instructions[p.currentIndex][2]]
		} else {
			value = val
		}
	}

	switch command {
		case "set":
			p.register[l] = value
		case "add":
			p.register[l] += value
		case "mul": // multiply
			p.register[l] *= value
		case "mod": // modulo

			p.register[l] %= value
		case "rcv": // receive
			if len(*receivedValues) > 0 {
				p.isWaiting = false
				p.register[l] = (*receivedValues)[0]
				*receivedValues = (*receivedValues)[1:]
			} else {
				p.isWaiting = true
			}
		case "jgz": // jumps
			if v, err := strconv.Atoi(l); err == nil {
				if v > 0 {
					p.currentIndex += value - 1
				}
			} else if p.register[l] > 0 {
				p.currentIndex += value - 1
			}
		case "snd": // send a value
			p.sentValAmount++
			p.sentValues = append(p.sentValues, p.register[l])
	} // switch

	if !p.isWaiting {
		p.currentIndex++
	}
}

func answerPart2 (instructions map[int][]string) {
	prog0 := Program{
		register:   make(map[string]int),
		sentValues: make([]int, 0),
	}
	prog1 := Program{
		register:   make(map[string]int),
		sentValues: make([]int,0),
	}
	prog0.initProgram(0)
	prog1.initProgram(1)


	for {
		prog0.execute(&prog1.sentValues, instructions)
		prog1.execute(&prog0.sentValues, instructions)
		if prog0.isWaiting && prog1.isWaiting {
			break
		}
	}
	fmt.Println("Answer part 2:", prog1.sentValAmount)

}
func main() {
	in := getInput()
	answerPart1(in)
	answerPart2(in)
}
// 1667745 to high
// 254 to low
