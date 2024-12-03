package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	_, err = file.Seek(0, 0)
	if err != nil {
		panic(fmt.Sprintf("Failed to reset file position: %v", err))
	}

	validDataWithConditionals := ExtractValidDataWithConditionals(file)
	multiplicationsOutputWithConditionals := CalcMultiplications(validDataWithConditionals)
	println("Multiplications result w/ conditionals: ", multiplicationsOutputWithConditionals)
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

func ExtractValidDataWithConditionals(corruptedData io.Reader) [][]int {
	regexPattern := `mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\)`
	regex := regexp.MustCompile(regexPattern)

	validMultiplications := make([][]int, 0)
	scanner := bufio.NewScanner(corruptedData)

	instructionsEnabled := true
	for scanner.Scan() {
		line := scanner.Text()

		matches := regex.FindAllSubmatch([]byte(line), -1)

		for _, match := range matches {
			if instructionsEnabled && strings.Contains(string(match[0]), "mul") {
				x, errX := strconv.Atoi(string(match[1]))
				if errX != nil {
					fmt.Printf("Could not convert %s into int: %v\n", match[1], errX)
					continue
				}

				y, errY := strconv.Atoi(string(match[2]))
				if errY != nil {
					fmt.Printf("Could not convert %s into int: %v\n", match[2], errY)
					continue
				}

				validMultiplications = append(validMultiplications, []int{x, y})
			} else {
				instruction := string(match[0])

				if instruction == "don't()" {
					instructionsEnabled = false
				} else if instruction == "do()" {
					instructionsEnabled = true
				}
			}
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
