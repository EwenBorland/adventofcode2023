package day_3_test

import (
	"aoc2023/day_3"
	"bufio"
	"os"
	"reflect"
	"testing"
)

var test_answer int = 4361

func Test_Day_3(t *testing.T) {
	f, err := os.Open("mock_input.txt")
	if err != nil{
		t.Fatalf("file io failed, err: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	output := day_3.Day_3(scanner)
	if output != test_answer {
		t.Errorf("Day(3) fault, want: %v, got: %v", test_answer, output)
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

func TestEvaluatePart(t *testing.T) {
	testCases := []struct {
		name            string
		x               int
		line            string
		expectedValue   string
		expectedIndices []int
	}{
		{"123.", 0, "123.", "123", []int{0, 1, 2}},
		{".123.", 1, ".123.", "123", []int{1, 2, 3}},
		{".123$", 1, ".123$", "123", []int{1, 2, 3}},
		{"..1", 2, "..1", "1", []int{2}},
		{"..1.1", 2, "..1.1", "1", []int{2}},
		{"..1..1", 5, "..1..1", "1", []int{5}},
		{"..1..12", 5, "..1..12", "12", []int{5, 6}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := []day_3.Part{}
			day_3.EvaluatePart(tc.x, 0, tc.line, &p)
			if len(p) == 0 {
				t.Fatalf("TestEvaluatePart: no part added to list")
			}
			if tc.expectedValue != p[0].Value {
				t.Errorf("TestEvaluatePart: Incorrect response value for line '%v', expected %v, got %v", tc.line, tc.expectedValue, p[0].Value)
			}
			if !reflect.DeepEqual(tc.expectedIndices, p[0].Indices){
				t.Errorf("TestEvaluatePart: Incorrect indices for line '%v', expected %v, got %v", tc.line, tc.expectedIndices, p[0].Indices)
			}
		})
	}
}

func TestEvaluateSymbol(t *testing.T) {
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
			day_3.EvaluateSymbol(tc.x, tc.y, &m)

			if !reflect.DeepEqual(tc.expectedIndices, m) {
				t.Errorf("TestEvaluateSymbol: Incorrect map , expected %v, got %v", tc.expectedIndices, m)
			}
		})
	}
}

func TestParseLine(t *testing.T){
	t.Run("TestParseLine", func(t *testing.T) {
		p := []day_3.Part{}
		m := map[int][]int{}
		day_3.ParseLine("..1$.234", 5, &p, &m)
		expectedP := []day_3.Part{
			{Line: 5, Value: "1", Indices: []int{2}},
			{Line: 5, Value: "234", Indices: []int{5,6,7}},
		}
		expectedM := map[int][]int{
			4:{2,3,4},
			5:{2,3,4},
			6:{2,3,4},
		}

		if !reflect.DeepEqual(p, expectedP){
			t.Errorf("Incorrect parts , expected %v, got %v", expectedP, p)
		}
		if !reflect.DeepEqual(m, expectedM){
			t.Errorf("Incorrect map , expected %v, got %v", expectedM, m)
		}
	})
}
