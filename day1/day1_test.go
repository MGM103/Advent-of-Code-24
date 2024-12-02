package main_test

import (
	"reflect"
	"strings"
	"testing"

	day1 "github.com/mgm103/advent-of-code-24/day1"
)

func TestReadDistances(t *testing.T) {
	mockDistances := ` 3 4
		4 3
		2 5
		1 3
		3 9
		3 3`

	mockDistancesFile := strings.NewReader(mockDistances)

	t.Run("reads distances into slices", func(t *testing.T) {

		listA, listB := day1.ReadDistances(mockDistancesFile)
		expectedListA := []int{1, 2, 3, 3, 3, 4}
		expectedListB := []int{3, 3, 3, 4, 5, 9}

		if len(listA) != len(expectedListA) {
			t.Errorf("Expected length of: %d, instead got %d", len(expectedListA), len(listA))
		}

		if len(listB) != len(expectedListB) {
			t.Errorf("expected length of: %d, instead got %d", len(expectedListB), len(listB))
		}

		if !reflect.DeepEqual(listA, expectedListA) {
			t.Errorf("Expected %v, instead got %v", expectedListA, listA)
		}

		if !reflect.DeepEqual(listB, expectedListB) {
			t.Errorf("Expected %v, instead got %v", expectedListB, listB)
		}
	})
}

func TestCalculateTotalDistance(t *testing.T) {
	t.Run("outputs sum of differences", func(t *testing.T) {
		distances1 := []int{1, 2, 3, 3, 3, 4}
		distances2 := []int{3, 3, 3, 4, 5, 9}

		totalDistance := day1.CalculateTotalDistance(distances1, distances2)
		expectedDistance := 11

		if totalDistance != expectedDistance {
			t.Errorf("Expected %d, instead got %d", expectedDistance, totalDistance)
		}
	})
}

func TestCalculateSimilarityScore(t *testing.T) {
	t.Run("calculates similarity score", func(t *testing.T) {
		distances1 := []int{1, 2, 3, 3, 3, 4}
		distances2 := []int{3, 3, 3, 4, 5, 9}

		totalDistance := day1.CalculateSimilarityScore(distances1, distances2)
		expectedDistance := 31

		if totalDistance != expectedDistance {
			t.Errorf("Expected %d, instead got %d", expectedDistance, totalDistance)
		}
	})
}
