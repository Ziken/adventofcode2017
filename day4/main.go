package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"sort"
)

const INPUT_FILE = "input.txt"

func check (e error) {
	if e != nil {
		panic(e)
	}
}
func getInput() (in [][]string) {
	file, errFile := os.Open(INPUT_FILE)
	check(errFile)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	check(scanner.Err())
	for scanner.Scan() {
		row := scanner.Text()
		in = append(in, strings.Split(row, " "))
	}
	return in
}
func areAnagrams(w1,w2 string) bool {
	if len(w1) != len(w2) {
		return false
	}

	s1 := strings.Split(w1,"")
	sort.Strings(s1)

	s2 := strings.Split(w2,"")
	sort.Strings(s2)
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true

}
func answerPart1 (passphrases [][]string) {
	var validPassphrasesAmount int
	for _,row := range passphrases {
		ok := true
		for i,pha := range row {
			for j := i+1; j < len(row); j++ {
				if pha == row[j] {
					ok = false
				}
			}
		}
		if ok {
			validPassphrasesAmount++
		}
	}

	fmt.Println("Answer part 1:", validPassphrasesAmount)
}
func answerPart2 (passphrases [][]string) {
	var validPassphrasesAmount int
	for _,row := range passphrases {
		ok := true
		for i,pha := range row {
			for j := i+1; j < len(row) && ok; j++ {
				if areAnagrams(pha, row[j]) {
					ok = false
				}
			}
		}
		if ok {
			validPassphrasesAmount++
		}
	}

	fmt.Println("Answer part 2:", validPassphrasesAmount)
}
func main() {
	in := getInput()
	answerPart1(in)
	answerPart2(in)
}
