package day_3_test

import (
	"aoc2023/day_3"
	"aoc2023/helpers"
	"testing"
)

var test_answer int = 4361

func Test_Day_3(t *testing.T) {
	scanner, err := helpers.LoadTestScanner("mock_input.txt")
	if err != nil{
		t.Fatalf("file io failed, err: %v", err)
	}
	
	output := day_3.Day_3(scanner)
	if output != test_answer {
		t.Errorf("Day(3) fault, want: %v, got: %v", test_answer, output)
	}
}
