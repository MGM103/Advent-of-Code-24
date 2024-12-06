package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	manualUpdatesFile, err := os.Open("manualUpdates.txt")
	if err != nil {
		panic(fmt.Sprintf("Could not open file: %v\n", err))
	}
	defer manualUpdatesFile.Close()

	orderingData := GetOrderData(manualUpdatesFile)
	manualPageOrders := GetPageOrders(manualUpdatesFile)
	validOrderings, invalidOrderings := GetGroupedPageOrders(orderingData, manualPageOrders)
	correctedOrderings := ValidateIncorrectOrderings(invalidOrderings, orderingData)
	sumOfMiddlePages := GetSumMiddlePagesValidOrderings(validOrderings)
	sumOfCorrectedOrderings := GetSumMiddlePagesValidOrderings(correctedOrderings)
	println("Sum of middle pages: ", sumOfMiddlePages)
	println("Sum of corrected orderings middle pages", sumOfCorrectedOrderings)
}

func GetOrderData(manualUpdates io.ReadSeeker) map[int][]int {
	orderingData := make(map[int][]int, 0)
	scanner := bufio.NewScanner(manualUpdates)
	filePosition := 0

	for scanner.Scan() {
		lineContents := scanner.Text()
		filePosition += len(lineContents) + 1 // + 1 for \n

		// Exit when blank line encountered
		if strings.TrimSpace(lineContents) == "" {
			break
		}

		ordering := strings.Split(lineContents, "|")
		if len(ordering) != 2 {
			fmt.Println("Incorrect content in line: ", lineContents)
			continue
		}

		pageAppearsFirst, errAppearsFirst := strconv.Atoi(ordering[0])
		if errAppearsFirst != nil {
			fmt.Printf("Error reading first page number %s: %v", ordering[0], ordering)
			continue
		}
		pageAppearsAfter, errAppearsAfter := strconv.Atoi(ordering[1])
		if errAppearsAfter != nil {
			fmt.Printf("Error reading second page number %s: %v", ordering[0], ordering)
		}

		orderingData[pageAppearsFirst] = append(orderingData[pageAppearsFirst], pageAppearsAfter)
	}

	// Set file to page orders line
	manualUpdates.Seek(int64(filePosition), io.SeekStart)

	return orderingData
}

func GetPageOrders(manualUpdates io.ReadSeeker) [][]int {
	pageOrders := make([][]int, 0)
	scanner := bufio.NewScanner(manualUpdates)

	for scanner.Scan() {
		lineContents := scanner.Text()
		pageOrder := strings.Split(lineContents, ",")
		intPageOrder := make([]int, len(pageOrder))

		for i, pageNum := range pageOrder {
			intPageOrder[i], _ = strconv.Atoi(pageNum)
		}

		pageOrders = append(pageOrders, intPageOrder)
	}

	return pageOrders
}

func GetGroupedPageOrders(orderingData map[int][]int, pageOrders [][]int) ([][]int, [][]int) {
	validPageOrders := make([][]int, 0)
	invalidPageOrders := make([][]int, 0)

	for _, pageOrder := range pageOrders {
		pagesSeen := make(map[int]bool, 0)
		validOrder := true

		for _, pageNum := range pageOrder {
			pagesSeen[pageNum] = true

			for _, requiredPageAfter := range orderingData[pageNum] {
				if pagesSeen[requiredPageAfter] {
					validOrder = false
					break
				}
			}

			if !validOrder {
				break
			}
		}

		if validOrder {
			validPageOrders = append(validPageOrders, pageOrder)
		} else {
			invalidPageOrders = append(invalidPageOrders, pageOrder)
		}
	}

	return validPageOrders, invalidPageOrders
}

func ValidateIncorrectOrderings(invalidOrderings [][]int, orderingData map[int][]int) [][]int {
	for _, invalidOrder := range invalidOrderings {
		for i := range invalidOrder {
			for j := len(invalidOrder) - 1; j > i; j-- {
				for _, requiredPageAfter := range orderingData[invalidOrder[j]] {
					if requiredPageAfter == invalidOrder[i] {
						temp := invalidOrder[i]
						invalidOrder[i] = invalidOrder[j]
						invalidOrder[j] = temp
					}
				}
			}

		}
	}

	return invalidOrderings
}

func GetSumMiddlePagesValidOrderings(validOrderings [][]int) int {
	sumMiddlePages := 0

	for _, validOrdering := range validOrderings {
		middleIndex := len(validOrdering) / 2
		sumMiddlePages += validOrdering[middleIndex]
	}

	return sumMiddlePages
}
