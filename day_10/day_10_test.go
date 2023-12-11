package day_10_test

import (
	"aoc2023/day_10"
	"bufio"
	"os"
	"slices"
	"testing"
)

var test_answer_part_1 int = 8
var test_answer_part_2 int = 10

func Test_day_10(t *testing.T) {
	f, err := os.Open("mock_input.txt")
	if err != nil {
		t.Fatalf("file io failed, err: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	output_part1, output_part2 := day_10.Day_10(scanner)
	if output_part1 != test_answer_part_1 {
		t.Errorf("day_10 part 1 fault, want: %v, got: %v", test_answer_part_1, output_part1)
	}
	if output_part2 != test_answer_part_2 {
		t.Errorf("day_10 part 2 fault, want: %v, got: %v", test_answer_part_2, output_part2)
	}
}

func TestParseLine(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		outputLine  []rune
		outputStart int
	}{
		{"..F7.", "..F7.", []rune{'.', '.', 'F', '7', '.'}, -1},
		{"SJ.L7", "SJ.L7", []rune{'S', 'J', '.', 'L', '7'}, 0},
		{"|F--J", "|F--J", []rune{'|', 'F', '-', '-', 'J'}, -1},
		{"|FS-J", "|FS-J", []rune{'|', 'F', 'S', '-', 'J'}, 2},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			outputLine, outputStart := day_10.ParseLine(tc.input)
			if outputStart != tc.outputStart {
				t.Errorf("TestParseLine: Incorrect response for '%v', got: %v, expected: %v)", tc.name, outputStart, tc.outputStart)
			}
			if !slices.Equal(outputLine, tc.outputLine) {
				t.Errorf("TestParseLine: Incorrect response for '%v', got: %v, expected: %v)", tc.name, outputLine, tc.outputLine)
			}
		})
	}

}
