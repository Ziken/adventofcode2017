package main

import (
	"fmt"
	"io/ioutil"
)

const INPUT_FILE = "input.txt"
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func getInput() string {
	bytes, err := ioutil.ReadFile(INPUT_FILE)
	check(err)
	return string(bytes)
}
func answerPart1And2(streamOfCharacters string) {
	isGarbageOpen := false
	ignoreSign := false
	levelNesting := 0
	score := 0
	ignoredSignsAmount := 0
	for _, s := range streamOfCharacters {
		sign := string(s)
		switch {
			case ignoreSign:
				ignoreSign = false
			case sign == ">":
				isGarbageOpen = false
			case sign == "!" && isGarbageOpen:
				ignoreSign = true
			case isGarbageOpen:
				ignoredSignsAmount++
			case sign == "<":
				isGarbageOpen = true
			case sign == "{":
				levelNesting++
				score += levelNesting
			case sign == "}":
				levelNesting--
		}
	}
	fmt.Println("Answer part 1:", score)
	fmt.Println("Answer part 2:", ignoredSignsAmount)

}
func main() {
	in := getInput()
	answerPart1And2(in)
}