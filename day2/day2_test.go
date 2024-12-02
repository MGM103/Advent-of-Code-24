package main_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	day2 "github.com/mgm103/advent-of-code-24/day2"
)

func TestIsLevelsDecreasing(t *testing.T) {
	cases := []struct {
		name           string
		input          []int
		expectedOutput bool
	}{
		{"Is decreasing", []int{7, 6, 4, 2, 1}, true},
		{"Is not decreasing", []int{7, 7, 4, 2, 1}, false},
		{"Is increasing", []int{1, 2, 4, 6, 7}, false},
		{"Margin too wide", []int{9, 7, 6, 2, 1}, false},
		{"Is empty", []int{}, false},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			got := day2.IsLevelsDecreasing(test.input)

			if got != test.expectedOutput {
				t.Errorf("Expected %t, instead got: %t", test.expectedOutput, got)
			}
		})
	}
}

func TestIsLevelsIncreasing(t *testing.T) {
	cases := []struct {
		name           string
		input          []int
		expectedOutput bool
	}{
		{"Is increasing", []int{1, 3, 6, 7, 9}, true},
		{"Is not decreasing", []int{1, 3, 3, 6, 7}, false},
		{"Is decreasing", []int{7, 6, 4, 2, 1}, false},
		{"Margin too wide", []int{1, 2, 6, 7, 9}, false},
		{"Is empty", []int{}, false},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			got := day2.IsLevelsIncreasing(test.input)

			if got != test.expectedOutput {
				t.Errorf("Expected %t, instead got: %t", test.expectedOutput, got)
			}
		})
	}
}

func TestIsSafe(t *testing.T) {
	cases := []struct {
		name           string
		input          []int
		expectedOutput bool
	}{
		{"Is increasing", []int{1, 3, 6, 7, 9}, true},
		{"Is not increasing", []int{1, 3, 3, 6, 7}, false},
		{"Is decreasing", []int{7, 6, 4, 2, 1}, true},
		{"Is not decreasing", []int{7, 7, 4, 2, 1}, false},
		{"Increasing margin too wide", []int{1, 2, 6, 7, 9}, false},
		{"Decreasing Margin too wide", []int{9, 7, 6, 2, 1}, false},
		{"Is empty", []int{}, false},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			got := day2.IsSafe(test.input)

			if got != test.expectedOutput {
				t.Errorf("Expected %t, instead got: %t", test.expectedOutput, got)
			}
		})
	}
}

func TestStringSliceToIntSlice(t *testing.T) {
	stringSlice := []string{"7", "6", "4", "2", "1"}

	t.Run("Test conversion is correct", func(t *testing.T) {
		expectedIntSlice := []int{7, 6, 4, 2, 1}

		intSlice, err := day2.StringSliceToIntSlice(stringSlice)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if !reflect.DeepEqual(intSlice, expectedIntSlice) {
			t.Errorf("Expected %v, instead got: %v", expectedIntSlice, intSlice)
		}
	})
}

func TestGetSubReport(t *testing.T) {
	report := []int{1, 3, 2, 4, 5}
	want := [][]int{{3, 2, 4, 5}, {1, 2, 4, 5}, {1, 3, 4, 5}, {1, 3, 2, 5}, {1, 3, 2, 4}}

	for i, _ := range report {
		t.Run(fmt.Sprintf("Generates sub-report omitting index %d", i), func(t *testing.T) {
			got := day2.GetSubReport(i, report)

			if !reflect.DeepEqual(got, want[i]) {
				t.Errorf("Expected %v, instead got: %v", want[i], got)
			}
		})
	}
}

func TestCalculateNumSafeReports(t *testing.T) {
	reportsData := `7 6 4 2 1
	1 2 7 8 9
	9 7 6 2 1
	1 3 2 4 5
	8 6 4 4 1
	1 3 6 7 9`
	mockReportsData := strings.NewReader(reportsData)

	t.Run("Calculates the number of safe reports correctly", func(t *testing.T) {
		want := 2
		got := day2.CalculateNumSafeReports(mockReportsData)

		if got != want {
			t.Errorf("Expected %d, instead got: %d", want, got)
		}
	})
}

func TestCalculateNumSafeReportsWithDampener(t *testing.T) {
	reportsData := `65 68 71 72 71
31 34 36 37 37
80 83 84 86 87 90 92 96
30 33 36 39 45
21 22 25 23 24`

	mockReportsData := strings.NewReader(reportsData)

	t.Run("Calculates the number of safe reports correctly", func(t *testing.T) {
		want := 5
		got := day2.CalculateNumSafeReportsWithDampener(mockReportsData)

		if got != want {
			t.Errorf("Expected %d, instead got: %d", want, got)
		}
	})
}
