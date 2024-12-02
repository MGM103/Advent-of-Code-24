package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("distancesList.txt")

	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %v", err))
	}
	defer file.Close()

	distances1, distances2 := ReadDistances(file)
	totalDistance := CalculateTotalDistance(distances1, distances2)
	similarityScore := CalculateSimilarityScore(distances1, distances2)

	fmt.Println("Total distance: ", totalDistance)
	fmt.Println("Similarity score: ", similarityScore)
}

func ReadDistances(file io.Reader) ([]int, []int) {
	var listA, listB []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		distancesLine := scanner.Text()

		distances := strings.Fields(distancesLine)
		if len(distances) != 2 {
			// bad to use Errorf instead?
			fmt.Println("Invalid line: ", distancesLine)
			continue
		}

		distance1, errD1 := strconv.Atoi(distances[0])
		distance2, errD2 := strconv.Atoi(distances[1])
		if errD1 != nil || errD2 != nil {
			fmt.Println("Error converting string to int", distancesLine)
		}

		listA = append(listA, distance1)
		listB = append(listB, distance2)
	}

	sort.Ints(listA)
	sort.Ints(listB)

	return listA, listB
}

func CalculateTotalDistance(distances1, distances2 []int) int {
	if len(distances1) != len(distances2) {
		return -1
	}

	var totalDistance int

	for i := 0; i < len(distances1); i++ {
		distanceBetweenPts := abs(distances1[i] - distances2[i])
		totalDistance += distanceBetweenPts
	}

	return totalDistance
}

func CalculateSimilarityScore(distances1, distances2 []int) int {
	if len(distances1) != len(distances2) {
		fmt.Println("Similarity score inputs have mismatching lengths")
		return -1
	}

	var similarityScore int

	for _, distance1 := range distances1 {
		count := 0
		for _, distance2 := range distances2 {
			if distance1 == distance2 {
				count++
			} else if distance1 < distance2 {
				break
			}
		}
		similarityScore += abs(distance1 * count)
	}

	return similarityScore
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
