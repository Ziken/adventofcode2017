package main

import (
	"io/ioutil"
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
func getInput() (input []int) {
	bytes, err := ioutil.ReadFile(INPUT_FILE)
	check(err)
	splittedRow := strings.Split(string(bytes), ",")
	for _, val := range splittedRow {
		n, errAtoi := strconv.Atoi(val)
		check(errAtoi);
		input = append(input, n)
	}
	return
}
func getInputAsASCII() (input []int) {
	bytes, err := ioutil.ReadFile(INPUT_FILE)
	check(err)
	inputAsByte := []byte(bytes)

	for _, b := range inputAsByte {
		input = append(input, int(b))
	}

	return
}
func fillList(list * []int) {
	listCap := cap(*list)
	for i:=0; i < listCap; i++ {
		*list = append(*list, i)
	}
}
func riverseList(list []int) {
	length := len(list)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}
func getSubList(list []int, cp, size int) []int {
	var subList []int
	listSize := len(list)
	for i:=0; i <  size; i++ {
		if cp >= listSize {
			cp = 0
		}
		subList = append(subList, list[cp])
		cp++
	}
	return subList
}
func applyList(list []int, cp, size int) {
	listSize := len(list)
	subList := getSubList(list, cp,size)
	riverseList(subList)
	for i:=0; i <  size; i++ {
		if cp >= listSize {
			cp = 0
		}
		list[cp] = subList[i]
		cp++
	}
}
func getDenseHash(list []int) (denseHash string) {
	for i := 0; i < 16; i++ {
		hashNum := list[i*16]
		for j := i*16+1; j < (i+1)*16; j++ {
			hashNum = xorTwoNums(hashNum, list[j])
		}
		hexNum := strconv.FormatInt(int64(hashNum), 16)
		if len(hexNum) <= 1 {
			denseHash += "0" + string(hexNum)
		} else {
			denseHash += hexNum
		}
	}

	return denseHash
}
func xorTwoNums(n1, n2 int) (result int) {
	if n2 > n1 {
		n1,n2 = n2,n1
	}
	var binN2 string
	power := 1
	binN1 := strconv.FormatInt(int64(n1),2)
	binTmp := strconv.FormatInt(int64(n2),2)
	lenDiff := len(binN1)-len(binTmp)
	for i:=0; i < lenDiff; i++{
		binN2 += "0"
	}
	binN2 += binTmp

	for i := len(binN1)-1; i >= 0; i-- {
		if binN1[i] != binN2[i] {
			result += power
		}
		power *= 2
	}

	return result
}
func answerPart1 (in []int) {
	listSize := 256
	var currentPos, skipSize int
	list := make([]int, 0, listSize)
	fillList(&list)

	for _, addRange := range in {
		applyList(list, currentPos, addRange)
		currentPos = (currentPos + addRange + skipSize)%listSize
		skipSize++
	}
	fmt.Println("Answer part 1:", list[0] * list[1])
}
func answerPart2(in []int) {
	listSize := 256
	var currentPos, skipSize int
	list := make([]int, 0, listSize)
	fillList(&list)
	in = append(in, 17, 31, 73, 47, 23)
	for i:=0; i < 64; i++ {
		for _, addRange := range in {
			applyList(list, currentPos, addRange)
			currentPos = (currentPos + addRange + skipSize)%listSize
			skipSize++
		}
	}
	fmt.Println("Answer part 2:", getDenseHash(list))
}

func main() {
	in := getInput()
	inAscii := getInputAsASCII()
	answerPart1(in)
	answerPart2(inAscii)

}
