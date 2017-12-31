package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"strconv"
	"regexp"
)

const INPUT_FILE = "input.txt"

func check (e error) {
	if e != nil {
		panic(e)
	}
}
type TuringBluePrint struct {
	name    int
	ifN     []map[string]int
}
func getInput() (int, int, []TuringBluePrint) {
	var startFrom int
	var steps int
	var bluePrints []TuringBluePrint

	file, errFile := os.Open(INPUT_FILE)
	check(errFile)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	check(scanner.Err())

	if scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		startFrom = int( []byte(row[len(row)-2])[0] )
	}
	if scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")
		s, errAtoi := strconv.Atoi(row[len(row)-2])
		check(errAtoi)
		steps = s
	}

	whichInst := 0
	whichCond := 0
	bluePrintIndex := -1
	for scanner.Scan() {
		row := scanner.Text()

		switch whichInst {
			case 0: // skip empty line
			case 1: // get state
				bluePrints = append(bluePrints, TuringBluePrint{})
				bluePrintIndex++
				bluePrints[bluePrintIndex].name = getState(row)
			case 2: // get condition
				whichCond = getIntFromString(row)
				bluePrints[bluePrintIndex].ifN = append(bluePrints[bluePrintIndex].ifN, map[string]int{})
			case 3: // write value n
				writeVal := getIntFromString(row)
				bluePrints[bluePrintIndex].ifN[whichCond]["write"] = writeVal
			case 4: // move where
				moveRight := strings.Split(row, "right")
				if len(moveRight) > 1 {
					bluePrints[bluePrintIndex].ifN[whichCond]["move"] = 1
				} else {
					bluePrints[bluePrintIndex].ifN[whichCond]["move"] = -1
				}
			case 5: //continue with state ...
				bluePrints[bluePrintIndex].ifN[whichCond]["next"] = getState(row)
				if whichCond == 0 { // reset parsing condition
					whichInst = 1
				} else {
					whichInst = -1
				}
		}
		whichInst++
	}
	return startFrom, steps, bluePrints
}
func getState(row string) int {
	state := row[len(row) - 2]
	return int(state)
}
func getIntFromString(row string) int {
	numRegexp := regexp.MustCompile(`([0-9]+)`)
	strNum := numRegexp.FindString(row)
	n,errAtoi := strconv.Atoi(strNum)
	check(errAtoi)
	return n
}
func answerPart1(startFrom, steps int, bluePrints []TuringBluePrint) {
	size := 100000
	tape := make([]int, size)
	currentPos := size/2
	next := startFrom
	for i := 0; i < steps; i++ {
		index := next - 'A'
		if bluePrints[index].name != next {
			for a := 0; a < len(bluePrints); a++ {
				if bluePrints[a].name == next {
					index = a
					break
				}
			}
		}
		currentVal := tape[currentPos]
		tape[currentPos] = bluePrints[index].ifN[currentVal]["write"]
		currentPos += bluePrints[index].ifN[currentVal]["move"]
		next = bluePrints[index].ifN[currentVal]["next"]
	}
	checksum := 0
	for _,v := range tape {
		checksum += v
	}
	fmt.Println("Answer part 1:", checksum)
}
func answerPart2() {
	fmt.Println("Answer part 2:", "Click \"Rebbot the Printer\" :>")
}
func main() {
	startFrom, steps, bluePrints := getInput()
	answerPart1(startFrom, steps, bluePrints)
	answerPart2()
}
