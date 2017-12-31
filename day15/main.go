package main

import (
	"os"
	"bufio"
	"strconv"
	"fmt"
	"strings"
)

const INPUT_FILE = "input.txt"

func check (e error) {
	if e != nil {
		panic(e)
	}
}
func getInput() (valueGenA, valueGenB int) {
	var values []int
	file, errFile := os.Open(INPUT_FILE)
	check(errFile)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	check(scanner.Err())

	for scanner.Scan() {
		row := scanner.Text()
		splitedRow := strings.Split(row, " ")
		val, errAtoi := strconv.Atoi(splitedRow[4])
		check(errAtoi)
		values = append(values, val)
	}
	valueGenA,valueGenB = values[0], values[1]
	return
}
func createGenerator(startValue, factor, multiples int) func () int {
	value := startValue
	divider := 2147483647

	return func () int {
		value = (value * factor) % divider
		for value % multiples != 0 {
			value = (value * factor) % divider
		}
		return value
	}
}
func replenishToNCharacters(intAsBin string, n int) (properNotation string) {
	properNotation = intAsBin
	if len(intAsBin) < n {
		amountOfLackZeros := n-len(intAsBin)
		zeros := ""
		for i := 0; i < amountOfLackZeros; i++ {
			zeros += "0"
		}
		properNotation = zeros + intAsBin
	}
	return
}
func compareFirst16Bits(valA,valB int) (isEqual bool) {
	aAsBin := replenishToNCharacters(strconv.FormatInt(int64(valA), 2), 16)
	bAsBin := replenishToNCharacters(strconv.FormatInt(int64(valB), 2), 16)
	aAsBin = aAsBin[len(aAsBin)-16:]
	bAsBin = bAsBin[len(bAsBin)-16:]
	for i:=0; i < 16; i++ {
		if aAsBin[i] != bAsBin[i] {
			return false
		}
	}
	return true

}
func answerPart1(startValGenA, startValGenB int) {
	factorGenA := 16807
	factorGenB := 48271
	genA, genB := createGenerator(startValGenA, factorGenA,1), createGenerator(startValGenB, factorGenB,1)
	amountOfComparisons := 40000000
	counter := 0
	for i := 0; i < amountOfComparisons; i++ {
		valA,valB := genA(),genB()
		if compareFirst16Bits(valA, valB) {
			counter++
		}
	}
	fmt.Println("Answer part 1: ", counter)
}
func answerPart2(startValGenA, startValGenB int) {
	factorGenA := 16807
	factorGenB := 48271
	genA, genB := createGenerator(startValGenA, factorGenA,4), createGenerator(startValGenB, factorGenB,8)
	amountOfComparisons := 5000000
	counter := 0
	for i := 0; i < amountOfComparisons; i++ {
		valA,valB := genA(),genB()
		if compareFirst16Bits(valA, valB) {
			counter++
		}

	}
	fmt.Println("Answer part 2: ", counter)
}
func main() {
	fmt.Println("It might take a while :)")
	a,b := getInput()
	answerPart1(a,b)
	answerPart2(a,b)
}
