package main

import (
	"os"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"math"
)

const INPUT_FILE = "input.txt"

func check (e error) {
	if e != nil {
		panic(e)
	}
}
func getInput() ([]map[string]int, []map[string]int, []map[string]int) {

	file, errFile := os.Open(INPUT_FILE)
	position     :=  make([]map[string]int, 0)
	velocity 	 :=  make([]map[string]int, 0)
	acceleration :=  make([]map[string]int, 0)

	check(errFile)
	scanner := bufio.NewScanner(file)
	check(scanner.Err())

	matchNumbers := regexp.MustCompile(`([\-0-9]+)`)
	for scanner.Scan() {
		row := scanner.Text()
		allIntsAsString := matchNumbers.FindAllString(row, 9)
		part := 0
		x,y,z := getXYZasInt(allIntsAsString[part],allIntsAsString[part+1], allIntsAsString[part+2])
		position = append(position, map[string]int{
			"x":x,
			"y":y,
			"z":z,
		})

		part += 3
		x,y,z = getXYZasInt(allIntsAsString[part],allIntsAsString[part+1], allIntsAsString[part+2])
		velocity = append(velocity, map[string]int{
			"x":x,
			"y":y,
			"z":z,
		})

		part += 3
		x,y,z = getXYZasInt(allIntsAsString[part],allIntsAsString[part+1], allIntsAsString[part+2])
		acceleration = append(acceleration, map[string]int{
			"x":x,
			"y":y,
			"z":z,
		})
	}
	return position, velocity, acceleration
}
func getXYZasInt(x,y,z string) (int,int,int) {
	ix, errX := strconv.Atoi(x)
	check(errX)
	iy, errY := strconv.Atoi(y)
	check(errY)
	iz, errZ := strconv.Atoi(z)
	check(errZ)

	return ix, iy, iz
}
func getDistance(coords map[string]int) int {
	abs := func (a int) int {
		if a < 0 {
			return  a * -1
		}
		return a
	}
	return abs(coords["x"]) + abs(coords["y"]) + abs(coords["z"])
}
func updatePosition(position, velocity, acceleration []map[string]int, i int) {
	velocity[i]["x"] += acceleration[i]["x"]
	velocity[i]["y"] += acceleration[i]["y"]
	velocity[i]["z"] += acceleration[i]["z"]

	position[i]["x"] += velocity[i]["x"]
	position[i]["y"] += velocity[i]["y"]
	position[i]["z"] += velocity[i]["z"]
}
func answerPart1(position, velocity, acceleration []map[string]int) {
	lastDistance := make([]int, len(position))
	var closestDistance = math.MaxInt32

	idClosetsParticle := 0

	for moves := 0; moves < 1000; moves++ {

		for i, coords := range position {
			updatePosition(position, velocity, acceleration, i)

			lastDistance[i] = getDistance(coords)
			if lastDistance[i] < lastDistance[idClosetsParticle] {
				idClosetsParticle = i
			}
		}
		if lastDistance[idClosetsParticle] < closestDistance {
			closestDistance = lastDistance[idClosetsParticle]
		}
	}
	fmt.Println("Answer part 1:", idClosetsParticle)
}
func isInSamePosition (pos1,pos2 map[string]int) bool {
	return pos1["x"] == pos2["x"] && pos1["y"] == pos2["y"] && pos1["z"] == pos2["z"]
}
func answerPart2(position, velocity, acceleration []map[string]int) {
	isColliding := make([]bool, len(position))
	// 100 moves is enough
	for moves := 0; moves < 100; moves++ {
		for a:=0; a < len(position); a++ {
			if isColliding[a] {
				continue
			}
			for b := a + 1; b < len(position); b++ {
				if isInSamePosition(position[a],position[b]) {
					isColliding[a] = true
					isColliding[b] = true
				}
			}
		}
		for i := range position {
			if !isColliding[i] {
				updatePosition(position, velocity, acceleration, i)
			}
		}
	}
	sum := 0
	for i:=0; i < len(isColliding); i++ {
		if !isColliding[i] {
			sum++
		}
	}
	fmt.Println("Answer Part 2:", sum)
}
func main() {
	pos, vel, acc := getInput()
	answerPart1(pos, vel, acc)
	pos1, vel1, acc1 := getInput()
	answerPart2(pos1, vel1, acc1)
}