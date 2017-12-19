package main

import (
	"io/ioutil"
	"strconv"
	"fmt"
)

const INPUT_FILE = "input.txt"
func check (e error) {
	if e != nil {
		panic(e)
	}
}
func getInput() int {
	dat, err := ioutil.ReadFile(INPUT_FILE)
	check(err)
	in, errAtoi := strconv.Atoi(string(dat))
	check(errAtoi)
	return in
}
func getSpinLock(steps, n int) []int{
	buff := make([]int,1,n)
	currentPos := 0
	currentVal := 0
	for i := 0; i < n; i++ {
		buffLen := len(buff)
		currentPos = (currentPos + steps) % buffLen
		buff = append(buff, 0)
		buffLen++
		currentVal++
		currentPos = (currentPos + 1) % buffLen
		prev := buff[currentPos]
		buff[currentPos] = currentVal
		// move elements
		for i:=currentPos+1;i<buffLen;i++ {
			prev,buff[i] = buff[i],prev
		}
	}
	return buff
}
func getValueAfter(n int, buff [] int) int{
	for i, value := range buff {
		if value == n {
			return buff[(i+1)%len(buff)]
		}
	}
	return -1
}
func answerPart1 (steps int) {
	buff := getSpinLock(steps, 2017)
	fmt.Println("Answer part 1:", getValueAfter(2017, buff))
}
func answerPart2 (steps int) {
	var currentPos, currentVal, answer int
	buffLen := 1
	amount := 50000000

	for i := 0; i < amount; i++ {

		currentPos = (currentPos + steps) % buffLen
		buffLen++
		currentVal++
		currentPos = (currentPos + 1) % buffLen
		if currentPos == 1 {
			answer = currentVal
		}
	}
	fmt.Println("Answer part 2:", answer)
}
func main() {
	//in := getInput()
	in := 355
	answerPart1(in)
	answerPart2(in)
}
/**
[0 8 10 2 1 4 5 9 7 3 6]			10
[0 100 76 28 48 39 46 42			100
[0 891 886 100 270 506				1000
[0 7027 4647 1731 4295 2096 3372	10000
[0 77204 83766 55086 49921			100000



 */