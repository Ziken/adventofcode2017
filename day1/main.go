package main

import (
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
)

const INPUT_FILE = "input.txt"

func check (e error) {
	if e != nil {
		panic(e)
	}
}
func getInput() []int {
	input := make([]int, 0)
	dat, err := ioutil.ReadFile(INPUT_FILE)
	check(err)
	splittedRow := strings.Split(string(dat), "")
	for _, val := range splittedRow {
		n, errAtoi := strconv.Atoi(val)
		check(errAtoi)
		input = append(input, n)
	}
	return input
}

func answerPart1(numbers []int) {
	captcha := 0
	size := len(numbers)
	for i,v := range numbers {
		if v == numbers[(i+1)%size] {
			captcha += v
		}
	}
	fmt.Println("Answer part 1:", captcha)
}
func answerPart2(numbers []int) {
	captcha := 0
	size := len(numbers)
	halfway := size/2
	for i,v := range numbers {
		if v == numbers[(i+halfway)%size] {
			captcha += v
		}
	}
	fmt.Println("Answer part 2:", captcha)
}
func main() {
	in := getInput()
	answerPart1(in)
	answerPart2(in)
}
