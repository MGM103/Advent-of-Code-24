package main_test

import (
	"reflect"
	"strings"
	"testing"

	day6 "github.com/mgm103/advent-of-code-24/day6"
)

func TestGetMapCoordinates(t *testing.T) {
	labMapContents := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`
	mockLabMapFile := strings.NewReader(labMapContents)
	wantLabMapLayout := [][]rune{
		{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '#', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '#', '.', '.', '^', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
	}
	wantStartingPos := day6.Coord{4, 6}
	gotLabMapLayout, gotStartingPos := day6.GetMapCoordinates(mockLabMapFile)

	if !reflect.DeepEqual(*gotStartingPos, wantStartingPos) {
		t.Errorf("Expected starting position cordinates of: %v, instead got: %v", wantStartingPos, *gotStartingPos)
	}

	if !reflect.DeepEqual(gotLabMapLayout, wantLabMapLayout) {
		t.Errorf("Expected map with the following layout: %v, instead got: %v", wantLabMapLayout, gotLabMapLayout)
	}
}

func TestCalcGuardPath(t *testing.T) {
	labLayout := [][]rune{
		{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '#', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '#', '.', '.', '^', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
	}
	startingPos := day6.Coord{4, 6}
	want := 41
	got := day6.CalcGuardPath(labLayout, startingPos)

	if got != want {
		t.Errorf("Expected %d, instead got: %d", want, got)
	}
}
