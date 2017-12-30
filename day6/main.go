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
func getInput() (in []int) {
	dat, err := ioutil.ReadFile(INPUT_FILE)
	check(err)
	intsAsStr := strings.Split(string(dat), "\t")
	for _, n := range intsAsStr {
		num, errAtoi := strconv.Atoi(n)
		check(errAtoi)
		in = append(in, num)
	}

	return
}
func answerPart1And2(in []int) {
	bankSize := len(in)
	memoryBanks := make([]int, bankSize)
	prevMemoryBanks := make(map[string]int)
	copy(memoryBanks, in)
	counter := 0
	for {
		i := findMax(memoryBanks)
		max := memoryBanks[i]
		memoryBanks[i] = 0

		for max > 0 {
			i = (i + 1) % bankSize
			memoryBanks[i]++
			max--
		}
		counter++

		arrStr := arrToString(memoryBanks)
		v := prevMemoryBanks[arrStr]

		if v != 0 {
			fmt.Println("Answer part 1:", len(prevMemoryBanks) + 1)
			fmt.Println("Answer part 2:", counter - v)
			break
		}
		prevMemoryBanks[arrStr] = counter
	}
}
func findMax(memoryBank []int) (indexMax int) {
	for i := 0; i < len(memoryBank); i++ {
		if memoryBank[i] > memoryBank[indexMax] {
			indexMax = i
		}
	}
	return
}
func arrToString (m []int) (s string){
	for _, value := range m {
		s += strconv.Itoa(value)
	}
	return
}
func main() {
	in := getInput()
	answerPart1And2(in)
}