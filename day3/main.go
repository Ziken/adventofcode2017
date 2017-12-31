package main

import (
	"io/ioutil"
	"strconv"
	"fmt"
	"math"
)

const INPUT_FILE = "input.txt"

func check (err error) {
	if err != nil {
		panic(err)
	}
}
func getInput () int {
	bytes, err  := ioutil.ReadFile(INPUT_FILE)
	check(err)
	n, errAtoi := strconv.Atoi(string(bytes))
	check(errAtoi)

	return n
}
func answerPart1 (n int){
	currentN := 1
	direction := 2
	length := 1 // length of sequence until turn
	x, y := 0,0
	foundNum := false
	for !foundNum {
		for t := 0; t < 2; t++ {
			for a := 0; a < length; a++ {
				currentN++
				switch direction {
					case 0: // left
						y--
					case 1: // down
						x++
					case 2: // right
						y++
					case 3: // up
						x--
				}
				if currentN == n {
					foundNum = true
					break
				}
			}
			if foundNum {
				break
			}
			direction = (direction+1)%4
		}
		length++
	}
	fmt.Println("Answer part 1:", math.Abs(float64(x)) + math.Abs(float64(y)))
}
func answerPart2(n int) {
	size := 20 // enough size
	var square [][]int
	for i := 0; i < size; i++ {
		square = append(square, make([]int, size))
	}
	x,y := size/2, size/2 // middle of array
	square[x][y] = 1
	direction := 2
	length := 1
	foundNum := false
	for !foundNum {
		for t := 0; t < 2; t++ {
			for a := 0; a < length; a++ {
				switch direction {
					case 0: // left
						y--
					case 1: // down
						x++
					case 2: // right
						y++
					case 3: // up
						x--
				}
				square[x][y] = square[x-1][y-1] + square[x-1][y] + square[x-1][y+1] +
							   square[x][y-1] + square[x][y] + square[x][y+1] +
							   square[x+1][y-1] + square[x+1][y] + square[x+1][y+1]
				if square[x][y] > n {
					foundNum = true
					break
				}
			}
			if foundNum {
				break
			}
			direction = (direction+1)%4
		}
		length++
	}

	fmt.Println("Answer part 2:", square[x][y])
}
func main() {
	in := getInput()
	answerPart1(in)
	answerPart2(in)
}
