package main

import (
	"io/ioutil"
	"fmt"
	"strconv"
)

const INPUT_FILE = "input.txt"

func check (e error) {
	if e != nil {
		panic(e)
	}
}
func getInputAsASCII() []int {
	bytes, err := ioutil.ReadFile(INPUT_FILE)
	check(err)

	return convertToAscii(bytes)
}
func convertToAscii(dat []byte) (input []int) {
	inputAsByte := dat

	for _, b := range inputAsByte {
		input = append(input, int(b))
	}
	return input
}
func fillList(list * []int) {
	listCap := cap(*list)
	for i:=0; i < listCap; i++ {
		*list = append(*list, i)
	}
}
func reverseList(list []int) {
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
	reverseList(subList)
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
			denseHash += "0" + hexNum
		} else {
			denseHash += hexNum
		}
	}
	return denseHash

}
func xorTwoNums(n1, n2 int) int{
	if n2 > n1 {
		n1,n2 = n2,n1
	}
	var result int
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
func createDenseHash(in []int) string{
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
	return getDenseHash(list)
}
func hexToBin(hex string) []byte {
	var binStr []byte
	for _,h := range hex {
		switch h {
		case 48:   // 0
			binStr = append(binStr, []byte("0000")...)
		case 49:   // 1
			binStr = append(binStr, []byte("0001")...)
		case 50:   // 2
			binStr = append(binStr, []byte("0010")...)
		case 51:   // 3
			binStr = append(binStr, []byte("0011")...)
		case 52:   // 4
			binStr = append(binStr, []byte("0100")...)
		case 53:   // 5
			binStr = append(binStr, []byte("0101")...)
		case 54:   // 6
			binStr = append(binStr, []byte("0110")...)
		case 55:   // 7
			binStr = append(binStr, []byte("0111")...)
		case 56:   // 8
			binStr = append(binStr, []byte("1000")...)
		case 57:   // 9
			binStr = append(binStr, []byte("1001")...)
		case 97:   // a
			binStr = append(binStr, []byte("1010")...)
		case 98:   // b
			binStr = append(binStr, []byte("1011")...)
		case 99:   // c
			binStr = append(binStr, []byte("1100")...)
		case 100:  // d
			binStr = append(binStr, []byte("1101")...)
		case 101:  // e
			binStr = append(binStr, []byte("1110")...)
		case 102:  // f
			binStr = append(binStr, []byte("1111")...)
		}
	}

	return binStr
}
func intToAscii(number int) (intAsAscii []int) {
	arrBytes := []byte(strconv.FormatInt(int64(number),10))
	for _, b := range arrBytes {
		intAsAscii = append(intAsAscii, int(b))
	}
	return

}
func createGird(in []int) (gird [][]byte) {
	for i:=0; i < 128; i++ {
		sequence := in
		sequence = append(sequence, int([]byte("-")[0]))
		sequence = append(sequence, intToAscii(i)...)
		gird = append(gird, hexToBin(createDenseHash(sequence)))
	}
	return
}
func answerPart1(gird [][]byte) {
	usedSquares := 0
	for _,row := range gird {
		for _, bit := range row {
			if bit == 49 {
				usedSquares++
			}
		}
	}
	fmt.Println("Answer part 1:", usedSquares)
}
func answerPart2(gird [][]byte) {
	amountOfRegions := 0
	// add additional row filled with 0 to prevent "panic: runtime error: index out of range"
	gird = append(gird, make([]byte,128))
	for i, row := range gird {
		for j, square := range row {
			if square == 49 {
				amountOfRegions++
				markRegion(gird, i,j)
			}
		}
	}
	fmt.Println("Answer part 2:", amountOfRegions)
}
func markRegion(gird [][]byte, i, j int) {
	cols := len(gird[0])
	for start := j; start < cols && gird[i][start] == 49; start++ {
		gird[i][start] = 0
		if gird[i+1][start] == 49 {
			markRegion(gird, i+1, start)
		}
		if i > 0 && gird[i-1][start] == 49 {
			markRegion(gird, i-1, start)
		}

	}
	start := j-1
	if start < 0 {
		start = 0
	}
	for ; start >= 0 && gird[i][start] == 49; start-- {
		gird[i][start] = 0
		if gird[i+1][start] == 49 {
			markRegion(gird, i+1, start)
		}
		if i > 0 && gird[i-1][start] == 49 {
			markRegion(gird, i-1, start)
		}
	}
}
func main() {
	inAscii := getInputAsASCII()
	gird := createGird(inAscii)
	answerPart1(gird)
	answerPart2(gird)
}