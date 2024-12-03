package main_test

import (
	"reflect"
	"strings"
	"testing"

	day3 "github.com/mgm103/advent-of-code-24/day3"
)

func TestExtractValidData(t *testing.T) {
	data := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	mockData := strings.NewReader(data)

	t.Run("extracts valid mul calls from data", func(t *testing.T) {
		got := day3.ExtractValidData(mockData)
		want := [][]int{{2, 4}, {5, 5}, {11, 8}, {8, 5}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, instead got %v", want, got)
		}
	})
}

func TestExtractValidDataWithConditionals(t *testing.T) {
	data := `where()(< }when()mul(678,62)%mul(747,584)from()-   select(59,725)mul(570,425)<~^:$where()!@}where()mul(542,816):don't()select()from()#(;!select()mul(541,668)mul(557,427)what()<#-why(952,537)*what()}mul(723,604)/^]'+;from()('-mul(227,146)'what()~*who():select()~from()do()mul(640,593)
	xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
	mockData := strings.NewReader(data)

	t.Run("extracts valid mul calls from data", func(t *testing.T) {
		got := day3.ExtractValidDataWithConditionals(mockData)
		want := [][]int{{678, 62}, {747, 584}, {570, 425}, {542, 816}, {640, 593}, {2, 4}, {8, 5}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, instead got %v", want, got)
		}
	})
}

func TestCalcMultiplications(t *testing.T) {
	cases := []struct {
		name  string
		input [][]int
		want  int
	}{
		{"Empty", make([][]int, 0), 0},
		{"Single", [][]int{{3, 2}}, 6},
		{"Multiple entries", [][]int{{2, 4}, {5, 5}, {11, 8}, {8, 5}}, 161},
		{"With conditionals", [][]int{{2, 4}, {8, 5}}, 48},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			got := day3.CalcMultiplications(test.input)

			if got != test.want {
				t.Errorf("Expected %d, instead got %d", test.want, got)
			}
		})
	}
}
