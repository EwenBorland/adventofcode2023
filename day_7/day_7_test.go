package day_7_test

import (
	"aoc2023/day_7"
	"bufio"
	"os"
	"testing"
)

var test_answer_part_1 int = 6440
var test_answer_part_2 int = 5905

func Test_day_7(t *testing.T) {
	f, err := os.Open("mock_input.txt")
	if err != nil {
		t.Fatalf("file io failed, err: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	output_part1, output_part2 := day_7.Day_7(scanner)
	if output_part1 != test_answer_part_1 {
		t.Errorf("day_7 part 1 fault, want: %v, got: %v", test_answer_part_1, output_part1)
	}
	if output_part2 != test_answer_part_2 {
		t.Errorf("day_7 part 2 fault, want: %v, got: %v", test_answer_part_2, output_part2)
	}
}

func TestGetCardMap(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output int
	}{
		{"5 of a kind", []int{1,1,1,1,1}, 1},
		{"4 of a kind:1", []int{1,1,1,1,2}, 2},
		{"4 of a kind:2", []int{3,1,1,1,1}, 2},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			handType := day_7.GetHandType(tc.input)
			if handType != tc.output {
				t.Errorf("TestGetCardMap: Incorrect response for '%v', got: %v, expected: %v)", tc.name, handType, tc.output)
			}
		})
	}

}
