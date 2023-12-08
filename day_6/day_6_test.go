package day_6_test

import (
	"aoc2023/day_6"
	"bufio"
	"os"
	"testing"
)

var test_answer_part_1 int = 288
var test_answer_part_2 int = 71503

func Test_Day_6(t *testing.T) {
	f, err := os.Open("mock_input.txt")
	if err != nil {
		t.Fatalf("file io failed, err: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	output_part1, output_part2 := day_6.Day_6(scanner)
	if output_part1 != test_answer_part_1 {
		t.Errorf("Day_6 part 1 fault, want: %v, got: %v", test_answer_part_1, output_part1)
	}
	if output_part2 != test_answer_part_2 {
		t.Errorf("Day_6 part 2 fault, want: %v, got: %v", test_answer_part_2, output_part2)
	}
}

func TestEvaluateRace(t *testing.T) {

	testCases := []struct {
		name      string
		inputTime int
		inputDist int
		options   int
	}{
		{"7 9", 7, 9, 4},
		{"15 40", 15, 40, 8},
		{"30 200", 30, 200, 9},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			options := day_6.EvaluateRace(tc.inputTime, tc.inputDist)
			if options != tc.options {
				t.Errorf("TestEvaluateRace: Incorrect response for '%v', expected %v, got %v", tc.name, tc.options, options)
			}
		})
	}

}
