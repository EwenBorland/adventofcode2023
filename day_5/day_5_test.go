package day_5_test

import (
	"aoc2023/day_5"
	"bufio"
	"os"
	"testing"
)

var test_answer_part_1 int64 = 46 //35
var test_answer_part_2 int64 = 0

func Test_Day_5(t *testing.T) {
	f, err := os.Open("mock_input.txt")
	if err != nil {
		t.Fatalf("file io failed, err: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	output_part1, output_part2 := day_5.Day_5(scanner)
	if output_part1 != test_answer_part_1 {
		t.Errorf("Day_5 part 1 fault, want: %v, got: %v", test_answer_part_1, output_part1)
	}
	if output_part2 != test_answer_part_2 {
		t.Errorf("Day_5 part 2 fault, want: %v, got: %v", test_answer_part_2, output_part2)
	}
}

// func TestRunExample(t *testing.T) {

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
