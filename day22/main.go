package main

import (
	"os"
	"bufio"
	"fmt"
)

const INPUT_FILE = "input.txt"

const INFECTED   =   byte(35) // #
const CLEAN      =   byte(46) // .
const WEAKENED   =   byte(87) // W
const FLAGGED    =   byte(70) // F
func check (e error) {
	if e != nil {
		panic(e)
	}
}
func getInput() (in [][]byte) {

	file, errFile := os.Open(INPUT_FILE)
	check(errFile)
	scanner := bufio.NewScanner(file)
	check(scanner.Err())
	for scanner.Scan() {
		row := scanner.Text()
		in = append(in, []byte(row))
	}

	return
}
func extendCluster(cluster [][]byte, space int) [][]byte {
	var newCluster [][]byte
	newSize := 2*space + len(cluster)
	for i := 0; i < space; i++ {
		newCluster = append(newCluster, make([]byte, newSize))
	}

	for a := 0; a < len(cluster); a++ {
		newCluster = append(newCluster, make([]byte, space))
		newCluster[a+space] = append(newCluster[a+space], cluster[a]...)
		newCluster[a+space] = append(newCluster[a+space], make([]byte, space)...)
	}

	for i := 0; i < space; i++ {
		newCluster = append(newCluster, make([]byte, newSize))
	}
	return newCluster
}
func getDirection () func(int,int,string) (int,int){
	/* Representation of array d (directions)
	   /|\
		1   0 ->
	 <- 0   0
	       \|/
	*/
	var d [][]int
	d = append(d, make([]int, 2))
	d = append(d, make([]int, 2))
	d[0][0] = 1
	return func (currentX, currentY int, direction string) (int, int) {
		switch direction {
			case "left":
				// rotate -90deg
				d[0][0], d[0][1], d[1][0], d[1][1] = d[1][0], d[0][0], d[1][1], d[0][1]
			case "right":
				// rotate +90deg
				d[0][0], d[0][1], d[1][0], d[1][1] = d[0][1], d[1][1], d[0][0], d[1][0]
			case "reverse":
				// double rotate -90deg => -180deg
				d[0][0], d[0][1], d[1][0], d[1][1] = d[1][0], d[0][0], d[1][1], d[0][1]
				d[0][0], d[0][1], d[1][0], d[1][1] = d[1][0], d[0][0], d[1][1], d[0][1]
		}
		
		up,left,right,down := -1*d[0][0],-1*d[0][1],d[1][0],d[1][1]
		currentX += up + down
		currentY += left + right
		return currentX, currentY
	}
}
func answerPart1(startingCluster [][]byte) {
	changeDirection := getDirection()
	var becameINFECTED int
	cluster := extendCluster(startingCluster, 1000)
	currentX := len(cluster)/2
	currentY := currentX
	for bursts := 0; bursts < 10000; bursts++ {
		node := cluster[currentX][currentY]
		
		if node == INFECTED { // INFECTED
			cluster[currentX][currentY] = CLEAN
			currentX,currentY = changeDirection(currentX, currentY, "right")
		} else if node == CLEAN || node == 0 { // CLEAN
			becameINFECTED++
			cluster[currentX][currentY] = INFECTED
			currentX,currentY = changeDirection(currentX, currentY, "left")
		}
		
	}
	fmt.Println("Answer part 1:", becameINFECTED)
}
func answerPart2(startingCluster [][]byte) {
	changeDirection := getDirection()
	cluster := extendCluster(startingCluster, 1000)

	var becameINFECTED int

	currentX := len(cluster)/2
	currentY := currentX
	for bursts := 0; bursts < 10000000; bursts++ {
		node := cluster[currentX][currentY]
		switch {
			case node == CLEAN || node == 0:
				cluster[currentX][currentY] = WEAKENED
				currentX,currentY = changeDirection(currentX, currentY, "left")
			case node == WEAKENED:
				becameINFECTED++
				cluster[currentX][currentY] = INFECTED
				// direction "same" means that a direction won't be changed
				currentX,currentY = changeDirection(currentX, currentY, "same")
			case node == INFECTED:
				cluster[currentX][currentY] = FLAGGED
				currentX,currentY = changeDirection(currentX, currentY, "right")
			case node == FLAGGED:
				cluster[currentX][currentY] = CLEAN
				currentX,currentY = changeDirection(currentX, currentY, "reverse")
		}
	}
	fmt.Println("Answer part 2:", becameINFECTED)
}
func main() {
	in := getInput()
	answerPart1(in)
	answerPart2(in)
}