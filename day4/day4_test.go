package main_test

import (
	"reflect"
	"strings"
	"testing"

	day4 "github.com/mgm103/advent-of-code-24/day4"
)

func TestLoadPuzzle(t *testing.T) {
	t.Run("Returns array with puzzle contents", func(t *testing.T) {
		puzzleText := `AMXFJA
ISAMXH
HABHAI
XMASUS
UXHGRU`
		mockPuzzle := strings.NewReader(puzzleText)
		want := [][]rune{{'A', 'M', 'X', 'F', 'J', 'A'}, {'I', 'S', 'A', 'M', 'X', 'H'}, {'H', 'A', 'B', 'H', 'A', 'I'}, {'X', 'M', 'A', 'S', 'U', 'S'}, {'U', 'X', 'H', 'G', 'R', 'U'}}
		got := day4.LoadPuzzle(mockPuzzle)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, instead got: %v", want, got)
		}
	})
}

func TestCalcHorizontal(t *testing.T) {
	puzzleText := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	mockPuzzleFile := strings.NewReader(puzzleText)
	mockPuzzle := day4.LoadPuzzle(mockPuzzleFile)

	t.Run("Returns occurrences of XMAS horizontally", func(t *testing.T) {
		want := 5
		got := day4.CalcHorizontalMatches(mockPuzzle, "XMAS")

		if got != want {
			t.Errorf("Expected %d, instead got: %d", want, got)
		}
	})
}

func TestCalcVertical(t *testing.T) {
	puzzleText := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	mockPuzzleFile := strings.NewReader(puzzleText)
	mockPuzzle := day4.LoadPuzzle(mockPuzzleFile)

	t.Run("Returns occurrences of XMAS vertically", func(t *testing.T) {
		want := 3
		got := day4.CalcVerticalMatches(mockPuzzle, "XMAS")

		if got != want {
			t.Errorf("Expected %d, instead got: %d", want, got)
		}
	})
}

func TestCalcDiagonal(t *testing.T) {
	puzzleText := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	mockPuzzleFile := strings.NewReader(puzzleText)
	mockPuzzle := day4.LoadPuzzle(mockPuzzleFile)

	t.Run("Returns occurrences of XMAS diagonally", func(t *testing.T) {
		want := 10
		got := day4.CalcDiagonalMatches(mockPuzzle, "XMAS")

		if got != want {
			t.Errorf("Expected %d, instead got: %d", want, got)
		}
	})
}

func TestCalcTotalMatches(t *testing.T) {
	puzzleText := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	mockPuzzleFile := strings.NewReader(puzzleText)
	mockPuzzle := day4.LoadPuzzle(mockPuzzleFile)

	t.Run("Returns total occurrences of XMAS", func(t *testing.T) {
		want := 18
		got := day4.CalcTotalMatches(mockPuzzle, "XMAS")

		if got != want {
			t.Errorf("Expected %d, instead got: %d", want, got)
		}
	})
}

func TestReverseWord(t *testing.T) {
	cases := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{"Even word", "Paranormal", "lamronaraP"},
		{"Odd word", "odd", "ddo"},
		{"Empty string", "", ""},
		{"Single character", "a", "a"},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			got := day4.ReverseString(test.input)

			if got != test.expectedOutput {
				t.Errorf("Expected %s, instead got: %s", test.expectedOutput, got)
			}
		})
	}
}
