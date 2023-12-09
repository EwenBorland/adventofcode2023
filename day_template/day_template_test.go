package day_template_test

import (
	"aoc2023/day_template"
	"bufio"
	"os"
	"testing"
)

var test_answer_part_1 int = 0
var test_answer_part_2 int = 0

func Test_day_template(t *testing.T) {
	f, err := os.Open("mock_input.txt")
	if err != nil {
		t.Fatalf("file io failed, err: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	output_part1, output_part2 := day_template.Day_template(scanner)
	if output_part1 != test_answer_part_1 {
		t.Errorf("day_template part 1 fault, want: %v, got: %v", test_answer_part_1, output_part1)
	}
	if output_part2 != test_answer_part_2 {
		t.Errorf("day_template part 2 fault, want: %v, got: %v", test_answer_part_2, output_part2)
	}
}

func TestRunExample(t *testing.T) {
	testCases := []struct {
		name   string
		input  int
		output bool
	}{
		{"1", 1, false},
		{"2", 2, true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isEven := tc.input%2 == 0
			if isEven != tc.output {
				t.Errorf("TestRunExample: Incorrect response for '%v', got: %v, expected: %v)", tc.name, isEven, tc.output)
			}
		})
	}

}
