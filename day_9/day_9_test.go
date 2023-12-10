package day_9_test

import (
	"aoc2023/day_9"
	"bufio"
	"os"
	"slices"
	"testing"
)

var test_answer_part_1 int64 = 114
var test_answer_part_2 int64 = 2

func Test_day_9(t *testing.T) {
	f, err := os.Open("mock_input.txt")
	if err != nil {
		t.Fatalf("file io failed, err: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	output_part1, output_part2 := day_9.Day_9(scanner)
	if output_part1 != test_answer_part_1 {
		t.Errorf("day_9 part 1 fault, want: %v, got: %v", test_answer_part_1, output_part1)
	}
	if output_part2 != test_answer_part_2 {
		t.Errorf("day_9 part 2 fault, want: %v, got: %v", test_answer_part_2, output_part2)
	}
}

func TestGetDifferences(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int64
		output []int64
	}{
		{"10  13  16  21  30  45  68", []int64{10, 13, 16, 21, 30, 45, 68}, []int64{3, 3, 5, 9, 15, 23}},
		{"1   3   6  10  15  21  28", []int64{1, 3, 6, 10, 15, 21, 28}, []int64{2, 3, 4, 5, 6, 7}},
		{"24 50...", []int64{24,50,87,141,223,341,487,621,653,420,-352,-2121,-5670,-12365}, []int64{26,37,54,82,118,146,134,32,-233,-772,-1769,-3549,-6695}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := day_9.GetDifferences(tc.input)
			if !slices.Equal(output, tc.output) {
				t.Errorf("TestGetDifferences: Incorrect response for '%v', got: %v, expected: %v)", tc.name, output, tc.output)
			}
		})
	}
}

func TestGetNextValue(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int64
		output int64
	}{
		{"10  13  16  21  30  45  68", []int64{10, 13, 16, 21, 30, 45}, 68},
		{"1   3   6  10  15  21  28", []int64{1, 3, 6, 10, 15, 21}, 28},
		{"24 50...", []int64{24,50,87,141,223,341,487,621,653,420,-352,-2121,-5670,-12365}, -24598},
		{"hmmm", []int64{1,0,-1,-2,-3,-4,-5}, -6},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := day_9.GetNextValue(tc.input)
			if output != tc.output {
				t.Errorf("TestGetNextValue: Incorrect response for '%v', got: %v, expected: %v)", tc.name, output, tc.output)
			}
		})
	}
}
