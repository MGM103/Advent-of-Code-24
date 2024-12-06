package main_test

import (
	"reflect"
	"strings"
	"testing"

	day5 "github.com/mgm103/advent-of-code-24/day5"
)

func TestGetOrderData(t *testing.T) {
	orderingText := `47|53
97|13
97|61
97|47
75|53`
	mockOrderingFile := strings.NewReader(orderingText)

	t.Run("Extracts correct ordering relationship from data", func(t *testing.T) {
		got := day5.GetOrderData(mockOrderingFile)
		want := map[int][]int{
			47: {53},
			97: {13, 61, 47},
			75: {53},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, instead got %v", want, got)
		}
	})
}

func TestGetPageOrders(t *testing.T) {
	pageOrdersText := `75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`
	mockPageOrdersFile := strings.NewReader(pageOrdersText)

	t.Run("Extracts page orders", func(t *testing.T) {
		got := day5.GetPageOrders(mockPageOrdersFile)
		want := [][]int{{75, 47, 61, 53, 29}, {97, 61, 53, 29, 13}, {75, 29, 13}, {75, 97, 47, 61, 53}, {61, 13, 29}, {97, 13, 75, 29, 47}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, instead got %v", want, got)
		}
	})
}

func TestGetValidPageOrders(t *testing.T) {
	manualText := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`
	mockManualFile := strings.NewReader(manualText)
	mockOrderData := day5.GetOrderData(mockManualFile)
	mockPageOrders := day5.GetPageOrders(mockManualFile)

	t.Run("Test valid page orders are retrieved", func(t *testing.T) {
		gotValid, gotInvalid := day5.GetGroupedPageOrders(mockOrderData, mockPageOrders)
		wantValid := [][]int{{75, 47, 61, 53, 29}, {97, 61, 53, 29, 13}, {75, 29, 13}}
		wantInvalid := [][]int{{75, 97, 47, 61, 53}, {61, 13, 29}, {97, 13, 75, 29, 47}}

		if !reflect.DeepEqual(gotValid, wantValid) {
			t.Errorf("Expected %v, instead got %v", wantValid, gotValid)
		}

		if !reflect.DeepEqual(gotInvalid, wantInvalid) {
			t.Errorf("Expected %v, instead got %v", wantInvalid, gotInvalid)
		}
	})
}

func TestValidateIncorrectOrderings(t *testing.T) {
	invalidOrderings := [][]int{{75, 97, 47, 61, 53}, {61, 13, 29}, {97, 13, 75, 29, 47}}
	orderingData := map[int][]int{
		47: {53, 13, 61, 29},
		97: {13, 61, 47, 29, 53, 75},
		75: {53, 47, 61, 13},
		29: {13},
		53: {29, 13},
		61: {53, 29},
	}

	t.Run("Modifies the order correctly", func(t *testing.T) {
		got := day5.ValidateIncorrectOrderings(invalidOrderings, orderingData)
		want := [][]int{{97, 75, 47, 61, 53}, {61, 29, 13}, {97, 75, 47, 29, 13}}
		gotSum := day5.GetSumMiddlePagesValidOrderings(got)
		wantSum := 123

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, instead got %v", want, got)
		}

		if gotSum != wantSum {
			t.Errorf("Expected %d, instead got: %v", wantSum, gotSum)
		}
	})
}

func TestGetSumMiddlePagesValidOrderings(t *testing.T) {
	validOrderings := [][]int{{75, 47, 61, 53, 29}, {97, 61, 53, 29, 13}, {75, 29, 13}}
	t.Run("Returns correct sum", func(t *testing.T) {
		got := day5.GetSumMiddlePagesValidOrderings(validOrderings)
		want := 143

		if got != want {
			t.Errorf("Expected %d, instead got %d", want, got)
		}
	})
}
