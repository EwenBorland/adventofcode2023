package day_12_test

import (
	"aoc2023/day_12"
	"bufio"
	"os"
	"testing"
)

var test_answer_part_1 int = 0
var test_answer_part_2 int = 0

func Test_day_12(t *testing.T) {
	f, err := os.Open("mock_input.txt")
	if err != nil {
		t.Fatalf("file io failed, err: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	output_part1, output_part2 := day_12.Day_12(scanner)
	if output_part1 != test_answer_part_1 {
		t.Errorf("day_12 part 1 fault, want: %v, got: %v", test_answer_part_1, output_part1)
	}
	if output_part2 != test_answer_part_2 {
		t.Errorf("day_12 part 2 fault, want: %v, got: %v", test_answer_part_2, output_part2)
	}
}

func TestCountArrangements(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output int
	}{
		{"1", "???.### 1,1,3", 1},
		{"2", ".??..??...?##. 1,1,3", 4},
		{"3", "?#?#?#?#?#?#?#? 1,3,1,6", 1},
		{"4", "????.#...#... 4,1,1", 1},
		{"5", "????.######..#####. 1,6,5", 4},
		{"6", "?###???????? 3,2,1", 10},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			count := day_12.CountArrangements(tc.input)
			if count != tc.output {
				t.Errorf("TestCountArrangements: Incorrect response for test %v, got: %v, expected: %v)", tc.name, count, tc.output)
			}
		})
	}

}
