package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("multiplicationData.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %v", err))
	}
	defer file.Close()

	validData := ExtractValidData(file)
	multiplicationsOutput := CalcMultiplications(validData)

	println("Multiplications result: ", multiplicationsOutput)
}

func ExtractValidData(corruptedData io.Reader) [][]int {
	regexPattern := `mul\((\d{1,3}),(\d{1,3})\)`
	regex := regexp.MustCompile(regexPattern)

	validMultiplications := make([][]int, 0)
	scanner := bufio.NewScanner(corruptedData)

	for scanner.Scan() {
		line := scanner.Text()

		matches := regex.FindAllSubmatch([]byte(line), -1)

		for _, match := range matches {
			x, errX := strconv.Atoi(string(match[1]))
			if errX != nil {
				fmt.Printf("Could not convert %s into int: %v", match[1], errX)
				continue
			}

			y, errY := strconv.Atoi(string(match[2]))
			if errY != nil {
				fmt.Printf("Could not convert %s into int: %v", match[2], errY)
				continue
			}

			validMultiplications = append(validMultiplications, []int{x, y})
		}
	}

	return validMultiplications
}

func CalcMultiplications(multiplications [][]int) int {
	total := 0

	if len(multiplications) == 0 {
		return total
	}

	for _, multiplication := range multiplications {
		product := multiplication[0] * multiplication[1]
		total += product
	}

	return total
}
