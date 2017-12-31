package main
import (
	"fmt"
	"strings"
	"io/ioutil"
	"math"
)

const INPUT_FILE = "input.txt"

func check (e error) {
	if e != nil {
		panic(e)
	}
}
func getInput() (input []string) {
	bytes, err := ioutil.ReadFile(INPUT_FILE)
	check(err)
	input = append(input, strings.Split(string(bytes), ",")...)
	return input
}
// https://www.redblobgames.com/grids/hexagons/
func answerPart1 (input []string) {
	var x,y,z float64
	for _, direction := range input {
		// cannot do it directly...
		a,b,c := preciseCoordinateBy(direction)
		x += a
		y += b
		z += c
	}
	fmt.Println("Distance:", (math.Abs(x)+math.Abs(y)+math.Abs(z))/2)
}
func preciseCoordinateBy(direction string) (x,y,z float64)  {
	switch direction {
		case "n":
		y += 1
		x += -1
		z += 0
		case "ne":
		y += 1
		x += 0
		z += -1
		case "se":
		y += 0
		x += 1
		z += -1
		case "s":
		y += -1
		x += 1
		z += 0
		case "sw":
		y += -1
		x += 0
		z += 1
		case "nw":
		y += 0
		x += -1
		z += 1
		default:
		panic("undefined")
	}
	return
}
func answerPart2(input []string) {
	var x, y, z, maxDistance float64
	for _, direction := range input {
		// cannot do it directly...
		a,b,c := preciseCoordinateBy(direction)
		x+=a
		y+=b
		z+=c
		currentDistance := (math.Abs(x)+math.Abs(y)+math.Abs(z))/2;
		if currentDistance > maxDistance {
			maxDistance = currentDistance
		}
	}
	fmt.Println("Max distance:", maxDistance)
}
func main() {
	in := getInput()
	answerPart1(in)
	answerPart2(in)
}
