package day_4_test

import (
	"aoc2023/day_4"
	"bufio"
	"os"
	"testing"
)

var test_answer int = 13

func Test_Day_4(t *testing.T) {
	f, err := os.Open("mock_input.txt")
	if err != nil{
		t.Fatalf("file io failed, err: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	output, _ := day_4.Day_4(scanner)
	if output != test_answer {
		t.Errorf("Day_4 fault, want: %v, got: %v", test_answer, output)
	}
}

// func TestIsSymbolOrPart(t *testing.T) {

// 	testCases := []struct {
// 		name     string
// 		input    rune
// 		isSymbol bool
// 		isPart   bool
// 	}{
// 		{"a", '1', false, true},
// 		{"c", 'c', true, false},
// 		{".", '.', false, false},
// 		{"$", '$', true, false},
// 		{"blank", ' ', false, false},
// 		{"%", '%', true, false},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			isSymbol, isPart := day_3.IsSymbolOrPart(tc.input)
// 			if isSymbol != tc.isSymbol || isPart != tc.isPart {
// 				t.Errorf("TestIsSymbolOrPart: Incorrect response for '%c', expected (%v, %v), got (%v, %v)", tc.input, tc.isSymbol, tc.isPart, isSymbol, isPart)
// 			}
// 		})
// 	}

// }