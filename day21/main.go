package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
)

const INPUT_FILE = "input.txt"

func check (e error) {
	if e != nil {
		panic(e)
	}
}
func getInput() (in []Pattern) {

	file, errFile := os.Open(INPUT_FILE)
	check(errFile)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	check(scanner.Err())
	indexRow := 0
	for scanner.Scan() {
		var pattern, transform [][]byte
		row := scanner.Text()
		splittedRow := strings.Split(row, " => ")
		rawPattern := strings.Split(splittedRow[0], "/")
		rawTransform := strings.Split(splittedRow[1], "/")
		for _, s := range rawPattern {
			pattern = append(pattern, []byte(s))
		}
		for _, s := range rawTransform {
			transform = append(transform, []byte(s))
		}
		in = append(in, Pattern{size:len(rawPattern), pattern:pattern, transformation:transform})
		indexRow++
	}

	return
}
type Pattern struct {
	pattern        [][]byte
	transformation [][]byte
	size           int
	idPermutation  int
}
func (p * Pattern) permute() {
	p.idPermutation++
	p.flip()
	if p.idPermutation % 2 == 0 {
		p.rotate()
	}
}
func (p * Pattern) rotate() {
	if p.size == 2 {
		t := make([]byte, 2)
		copy(t,p.pattern[1])
		p.pattern[1][1] = p.pattern[0][1]
		p.pattern[0][1] = p.pattern[0][0]

		p.pattern[0][0] = p.pattern[1][0]
		p.pattern[1][0] = t[1]

	} else {
		t := make([]byte, 3)
		copy(t, p.pattern[2])
		p.pattern[2][0] = p.pattern[2][2]
		p.pattern[2][1] = p.pattern[1][2]
		p.pattern[2][2] = p.pattern[0][2]

		p.pattern[0][2] = p.pattern[0][0]
		p.pattern[1][2] = p.pattern[0][1]


		p.pattern[0][0] = t[0]
		p.pattern[0][1] = p.pattern[1][0]

		p.pattern[1][0] = t[1]

	}
}
func (p * Pattern) flip() {
	if p.size == 2 {
		t := make([]byte, 2)
		copy(t,p.pattern[1])
		p.pattern[1][0] = p.pattern[0][0]
		p.pattern[1][1] = p.pattern[0][1]

		p.pattern[0][0] = t[0]
		p.pattern[0][1] = t[1]
	} else {
		t := make([]byte, 3)
		copy(t, p.pattern[2])
		p.pattern[2][0] = p.pattern[0][0]
		p.pattern[2][1] = p.pattern[0][1]
		p.pattern[2][2] = p.pattern[0][2]

		p.pattern[0][0] = t[0]
		p.pattern[0][1] = t[1]
		p.pattern[0][2] = t[2]
	}
}
func matchPattern(patterns []Pattern, part [][]byte) int {
	size := len(part)

	for i, p := range patterns {
		if p.size != size {
			continue
		}
		// loop for permute pattern
		for r := 0; r < 8; r++ {
			patternMatched := true
			for a := 0; a < size; a++ {
				for b := 0; b < size; b++ {
					if p.pattern[a][b] != part[a][b] {
						patternMatched = false
					}
				}
			}
			if patternMatched {
				return i
			}
			p.permute()
		}
	}
	return -1 // execute panic
}
func transformImage(image[][]byte, transformation [][]byte, i,j int) {
	size := len(transformation)

	for a := 0; a < size; a++ {
		for b := 0; b < size; b++ {
			image[i*size+a][j*size+b] = transformation[a][b]
		}
	}
}
func newEmptyImage(size int) (image [][]byte){
	for i := 0; i < size; i++ {
		image = append(image, make([]byte, size))
	}

	return
}
func getImageAfterNIterations(patterns [] Pattern, n int) [][]byte{
	var image [][]byte

	image = append(image, []byte(".#."))
	image = append(image, []byte("..#"))
	image = append(image, []byte("###"))

	for iterations := 0; iterations < n; iterations++ {
		imageSize := len(image)
		var partSize int
		if imageSize % 2 == 0 {
			partSize = 2
		} else {
			partSize = 3
		}

		newSize := imageSize/partSize * (partSize+1)
		newImage := newEmptyImage(newSize)

		for i := 0; i < imageSize/partSize; i++ {
			for j := 0; j < imageSize/partSize; j++ {
				part := newEmptyImage(partSize)
				// get part of the image
				for a := 0; a < partSize; a++{
					for b := 0; b < partSize; b++{
						part[a][b] = image[i*partSize+a][j*partSize+b]
					}
				}

				pIndex := matchPattern(patterns, part)
				transformImage(newImage, patterns[pIndex].transformation, i,j)
			}
		}
		image = newImage
	}
	return image
}
func countActivePixels(image [][]byte) (sum int) {
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image); j++ {
			if image[i][j] == 35 {
				sum++
			}
		}
	}
	return
}
func answerPart1(patterns []Pattern) {
	image := getImageAfterNIterations(patterns, 5)
	fmt.Println("Answer part 1:", countActivePixels(image))
}
func answerPart2(patterns []Pattern) {
	image := getImageAfterNIterations(patterns, 18)
	fmt.Println("Answer part 2:", countActivePixels(image))
}
func main() {
	in := getInput()
	answerPart1(in)
	answerPart2(in)
}
