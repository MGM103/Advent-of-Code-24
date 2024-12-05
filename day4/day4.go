package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	puzzleFile, err := os.Open("puzzle.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %v", err))
	}
	defer puzzleFile.Close()

	const WORD_TO_FIND string = "XMAS"

	puzzle := LoadPuzzle(puzzleFile)
	totalMatches := CalcTotalMatches(puzzle, WORD_TO_FIND)

	println("Total matches: ", totalMatches)
}

func LoadPuzzle(puzzleFile io.Reader) [][]rune {
	scanner := bufio.NewScanner(puzzleFile)
	puzzle := make([][]rune, 0)

	for scanner.Scan() {
		lineContents := scanner.Text()
		puzzleRow := []rune(lineContents)
		puzzle = append(puzzle, puzzleRow)
	}

	return puzzle
}

func CalcTotalMatches(puzzle [][]rune, word string) int {
	return CalcDiagonalMatches(puzzle, word) + CalcHorizontalMatches(puzzle, word) + CalcVerticalMatches(puzzle, word)
}

func CalcHorizontalMatches(puzzle [][]rune, word string) int {
	totalMatches := 0

	for _, row := range puzzle {
		rowStr := string(row)
		totalMatches += FindSubStringMatchesInString(rowStr, word)
	}

	return totalMatches
}

func CalcVerticalMatches(puzzle [][]rune, word string) int {
	totalMatches := 0

	for columnIndex := range puzzle[0] {
		var columnRunes []rune
		for rowIndex := range puzzle {
			columnRunes = append(columnRunes, puzzle[rowIndex][columnIndex])
		}
		colStr := string(columnRunes)
		totalMatches += FindSubStringMatchesInString(colStr, word)
	}

	return totalMatches
}

func CalcDiagonalMatches(puzzle [][]rune, word string) int {
	totalMatches := 0
	lenCol := len(puzzle[0])
	lenRow := len(puzzle)

	for col := range puzzle[0] {
		var diagonalRunes []rune

		for rowD, colD := 0, col; rowD < lenRow && colD < lenCol; rowD, colD = rowD+1, colD+1 {
			diagonalRunes = append(diagonalRunes, puzzle[rowD][colD])
		}

		diagonalStr := string(diagonalRunes)
		totalMatches += FindSubStringMatchesInString(diagonalStr, word)
	}

	for row := 1; row < lenRow; row++ {
		var diagonalRunes []rune
		var antiDiagonalRunes []rune

		for rowD, colD := row, 0; rowD < lenRow && colD < lenCol; rowD, colD = rowD+1, colD+1 {
			diagonalRunes = append(diagonalRunes, puzzle[rowD][colD])
		}

		for rowD, colD := row, lenCol-1; rowD < lenRow && colD >= 0; rowD, colD = rowD+1, colD-1 {
			antiDiagonalRunes = append(antiDiagonalRunes, puzzle[rowD][colD])
		}

		diagonalStr := string(diagonalRunes)
		antiDiagonalStr := string(antiDiagonalRunes)
		totalMatches += FindSubStringMatchesInString(diagonalStr, word) + FindSubStringMatchesInString(antiDiagonalStr, word)
	}

	for col := lenCol - 1; col > 0; col-- {
		var diagonalRunes []rune

		for rowD, colD := 0, col; rowD < lenRow && colD >= 0; rowD, colD = rowD+1, colD-1 {
			diagonalRunes = append(diagonalRunes, puzzle[rowD][colD])
		}

		diagonalStr := string(diagonalRunes)
		totalMatches += FindSubStringMatchesInString(diagonalStr, word)
	}

	return totalMatches
}

func FindSubStringMatchesInString(inputString, subString string) int {
	totalMatches := 0
	regexForward := regexp.MustCompile(subString)
	regexReverse := regexp.MustCompile(ReverseString(subString))

	forwardsMatches := regexForward.FindAllString(inputString, -1)
	reverseMatches := regexReverse.FindAllString(inputString, -1)
	totalMatches = len(forwardsMatches) + len(reverseMatches)

	return totalMatches
}

func ReverseString(word string) string {
	charWord := []rune(word)
	i, j := 0, len(charWord)-1

	for i < j {
		charWord[i], charWord[j] = charWord[j], charWord[i]
		i++
		j--
	}

	return string(charWord)
}
