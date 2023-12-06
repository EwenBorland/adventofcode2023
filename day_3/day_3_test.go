package day_3_test

import (
	"aoc2023/day_3"
	"bufio"
	"os"
	"reflect"
	"testing"
)

var test_answer1 int = 4361 // 522726
var test_answer2 int = 467835 // 

func Test_Day_3(t *testing.T) {
	f, err := os.Open("mock_input.txt")
	if err != nil {
		t.Fatalf("file io failed, err: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	output1, output2 := day_3.Day_3(scanner)
	if output1 != test_answer1 {
		t.Errorf("Day(3) fault 1, want: %v, got: %v", test_answer1, output1)
	}
	if output2 != test_answer2 {
		t.Errorf("Day(3) fault 2, want: %v, got: %v", test_answer2, output2)
	}
}

func TestIsSymbolOrPart(t *testing.T) {

	testCases := []struct {
		name     string
		input    rune
		isSymbol bool
		isPart   bool
	}{
		{"a", '1', false, true},
		{"c", 'c', true, false},
		{".", '.', false, false},
		{"$", '$', true, false},
		{"blank", ' ', false, false},
		{"%", '%', true, false},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isSymbol, isPart := day_3.IsSymbolOrPart(tc.input)
			if isSymbol != tc.isSymbol || isPart != tc.isPart {
				t.Errorf("TestIsSymbolOrPart: Incorrect response for '%c', expected (%v, %v), got (%v, %v)", tc.input, tc.isSymbol, tc.isPart, isSymbol, isPart)
			}
		})
	}

}

func TestParsePart(t *testing.T) {
	testCases := []struct {
		name            string
		x               int
		line            string
		expectedValue   int
		expectedIndices []int
	}{
		{"123.", 0, "123.", 123, []int{0, 1, 2}},
		{".123.", 1, ".123.", 123, []int{1, 2, 3}},
		{".123$", 1, ".123$", 123, []int{1, 2, 3}},
		{"..1", 2, "..1", 1, []int{2}},
		{"..1.1", 2, "..1.1", 1, []int{2}},
		{"..1..1", 5, "..1..1", 1, []int{5}},
		{"..1..12", 5, "..1..12", 12, []int{5, 6}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := []day_3.Part{}
			day_3.ParsePart(tc.x, 0, tc.line, &p)
			if len(p) == 0 {
				t.Fatalf("TestParsePart: no part added to list")
			}
			if tc.expectedValue != p[0].Value {
				t.Errorf("TestParsePart: Incorrect response value for line '%v', expected %v, got %v", tc.line, tc.expectedValue, p[0].Value)
			}
			if !reflect.DeepEqual(tc.expectedIndices, p[0].Indices) {
				t.Errorf("TestParsePart: Incorrect indices for line '%v', expected %v, got %v", tc.line, tc.expectedIndices, p[0].Indices)
			}
		})
	}
}

func TestParseSymbol(t *testing.T) {
	testCases := []struct {
		name            string
		x               int
		y               int
		expectedIndices map[int][]int
	}{
		{"middle indexes", 5, 7, map[int][]int{
			6: {4, 5, 6},
			7: {4, 5, 6},
			8: {4, 5, 6},
		}},
		{"left indexes", 0, 7, map[int][]int{
			6: {0, 1},
			7: {0, 1},
			8: {0, 1},
		}},
		{"top indexes", 5, 0, map[int][]int{
			0: {4, 5, 6},
			1: {4, 5, 6},
		}},
		{"left top indexes", 0, 0, map[int][]int{
			0: {0, 1},
			1: {0, 1},
		}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := map[int][]int{}
			s := []day_3.Symbol{}
			day_3.ParseSymbol(tc.x, tc.y, 'a', &s, &m)

			if !reflect.DeepEqual(tc.expectedIndices, m) {
				t.Errorf("TestParseSymbol: Incorrect map , expected %v, got %v", tc.expectedIndices, m)
			}
		})
	}
}

func TestParseLine(t *testing.T) {
	t.Run("TestParseLine", func(t *testing.T) {
		p := []day_3.Part{}
		s := []day_3.Symbol{}
		m := map[int][]int{}
		day_3.ParseLine("..1$.234", 5, &p, &s, &m)
		expectedP := []day_3.Part{
			{Line: 5, Value: 1, Indices: []int{2}},
			{Line: 5, Value: 234, Indices: []int{5, 6, 7}},
		}
		expectedM := map[int][]int{
			4: {2, 3, 4},
			5: {2, 3, 4},
			6: {2, 3, 4},
		}

		if !reflect.DeepEqual(p, expectedP) {
			t.Errorf("Incorrect parts , expected %v, got %v", expectedP, p)
		}
		if !reflect.DeepEqual(m, expectedM) {
			t.Errorf("Incorrect map , expected %v, got %v", expectedM, m)
		}
	})
}
