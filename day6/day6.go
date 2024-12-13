package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Coord struct {
	X int
	Y int
}

func (c *Coord) Update(x, y int) {
	c.X = x
	c.Y = y
}

func (c *Coord) IncrementX() {
	c.X++
}

func (c *Coord) IncrementY() {
	c.Y++
}

func (c *Coord) DecrementX() {
	c.X--
}

func (c *Coord) DecrementY() {
	c.Y--
}

func main() {
	mapFile, err := os.Open("labMap.txt")
	if err != nil {
		panic(fmt.Sprintf("Encountered error opening labMap.txt: %v", err))
	}
	defer mapFile.Close()

	labMapLayout, startingPos := GetMapCoordinates(mapFile)
	numDistinctPositions := CalcGuardPath(labMapLayout, *startingPos)
	fmt.Println("Number of distinct positions: ", numDistinctPositions)
}

func GetMapCoordinates(mapFile io.Reader) ([][]rune, *Coord) {
	mapLayout := make([][]rune, 0)
	startingPos := &Coord{0, 0}
	scanner := bufio.NewScanner(mapFile)
	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		yContents := make([]rune, len(line))

		for x, char := range line {
			yContents[x] = char

			if isSecurityGuard(char) {
				startingPos.Update(x, y)
			}
		}

		mapLayout = append(mapLayout, yContents)
		y++
	}

	return mapLayout, startingPos
}

func CalcGuardPath(labLayout [][]rune, startingPos Coord) int {
	visited := make(map[Coord]bool)
	numDistinctPos := 0
	lenX, lenY := len(labLayout[0]), len(labLayout)
	currentPos := startingPos
	visited[currentPos] = true

	for 0 <= currentPos.X && currentPos.X < lenX && 0 <= currentPos.Y && currentPos.Y < lenY {
		switch labLayout[currentPos.Y][currentPos.X] {
		case '^':
			if currentPos.Y-1 < 0 {
				currentPos.DecrementY()
			} else if labLayout[currentPos.Y-1][currentPos.X] == '#' {
				labLayout[currentPos.Y][currentPos.X] = '>'
			} else {
				labLayout[currentPos.Y][currentPos.X] = 'X'
				labLayout[currentPos.Y-1][currentPos.X] = '^'
				currentPos.DecrementY()
			}

		case '>':
			if currentPos.X+1 == lenX {
				currentPos.IncrementX()
			} else if labLayout[currentPos.Y][currentPos.X+1] == '#' {
				labLayout[currentPos.Y][currentPos.X] = 'v'
			} else {
				labLayout[currentPos.Y][currentPos.X] = 'X'
				labLayout[currentPos.Y][currentPos.X+1] = '>'
				currentPos.IncrementX()
			}
		case '<':
			if currentPos.X-1 < 0 {
				currentPos.DecrementX()
			} else if labLayout[currentPos.Y][currentPos.X-1] == '#' {
				labLayout[currentPos.Y][currentPos.X] = '^'
			} else {
				labLayout[currentPos.Y][currentPos.X] = 'X'
				labLayout[currentPos.Y][currentPos.X-1] = '<'
				currentPos.DecrementX()
			}
		case 'v':
			if currentPos.Y+1 == lenY {
				currentPos.IncrementY()
			} else if labLayout[currentPos.Y+1][currentPos.X] == '#' {
				labLayout[currentPos.Y][currentPos.X] = '<'
			} else {
				labLayout[currentPos.Y][currentPos.X] = 'X'
				labLayout[currentPos.Y+1][currentPos.X] = 'v'
				currentPos.IncrementY()
			}
		}

		if !visited[currentPos] {
			visited[currentPos] = true
			numDistinctPos++
		}
	}

	return numDistinctPos
}

func isSecurityGuard(char rune) bool {
	switch char {
	case '^':
		return true
	case '>':
		return true
	case '<':
		return true
	case 'v':
		return true
	default:
		return false
	}
}
