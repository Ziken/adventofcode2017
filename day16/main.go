package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"regexp"
	"strconv"
)

const INPUT_FILE = "input.txt"
func check (e error) {
	if e != nil {
		panic(e)
	}
}
func getInput() []string {
	dat, err := ioutil.ReadFile(INPUT_FILE)
	input := strings.Split(string(dat), ",")
	check(err)
	return input
}
func createPrograms() map[string]int {
	programs := make(map[string]int, 16)
	for i := 0; i < 16; i++ {
		programs[string(i+97)] = i
	}
	return programs
}
func getProgramsInOrder(programs map[string]int) string{
	order := make([]byte, 16)
	for key, pos := range programs {
		order[pos] = key[0]
	}
	return string(order)
}
func performDance(programs map[string]int, danceMoves []string) {
	for _, m := range danceMoves {
		executeMove(programs, m)
	}
}
func executeMove(programs map[string]int, move string) {
	switch move[0] {
	case 115: // s
		spinRegexp := regexp.MustCompile("([0-9]+)")
		spinNumber,_ := strconv.Atoi(spinRegexp.FindString(move))
		spin(programs, spinNumber)
	case 112: // p
		partner(programs, string(move[1]), string(move[3]))
	case 120: // x
		spinRegexp := regexp.MustCompile("([0-9]+)")
		found := spinRegexp.FindAll([]byte(move),2)
		p1,p2 := string(found[0]), string(found[1])
		pos1,_ := strconv.Atoi(p1)
		pos2,_ := strconv.Atoi(p2)
		exchange(programs, pos1,pos2)
	}
}
func spin(programs map[string]int, moves int) {
	pSize := len(programs)
	for key, pos := range programs {
		programs[key] = (pos+moves)%pSize
	}
}
func partner(programs map[string]int, p1,p2 string) {
	programs[p1],programs[p2] = programs[p2],programs[p1]
}
func exchange(programs map[string]int, pos1,pos2 int) {
	var key1, key2 string
	for prog, pos := range programs {
		if pos == pos1 {
			key1 = prog
		} else if pos == pos2 {
			key2 = prog
		}
	}
	partner(programs, key1, key2)
}
func answerPart1 (danceMoves []string) {
	programs := createPrograms()
	performDance(programs, danceMoves)

	fmt.Println("Answer part 1", getProgramsInOrder(programs))
}
func answerPart2(danceMoves []string) {
	programs := createPrograms()
	amountOfPerformances := 1000000000
	stop := false
	pastDances := make([]string,0)
	i := 0
	for ; i < amountOfPerformances; i++ {
		performDance(programs, danceMoves)
		pDance := getProgramsInOrder(programs)
		dLen := len(pastDances)-1
		for dLen >= 0 {
			if pastDances[dLen] == pDance {
				stop = true
				break
			}
			dLen--
		}
		if stop {
			break
		}
		pastDances = append(pastDances, pDance)
	}
	fmt.Println("Answer part 2", pastDances[amountOfPerformances % i - 1])

}
func main() {
	in := getInput()
	answerPart1(in)
	answerPart2(in)
}
