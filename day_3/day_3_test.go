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



func TestIsSymbolOrPart(t *testing.T){

	testCases := []struct{input rune; isSymbol bool; isPart bool}{
		{'1',false, true},
		{'c',true, false},
		{'.',false, false},
		{'$',true, false},
		{' ',false, false},
		{'%',true, false},

	}
	for _, tc := range testCases{
		isSymbol, isPart := day_3.IsSymbolOrPart(tc.input)
		if isSymbol != tc.isSymbol || isPart != tc.isPart{
			t.Errorf("TestIsSymbolOrPart: Incorrect response for '%c', expected (%v, %v), got (%v, %v)", tc.input, tc.isSymbol, tc.isPart, isSymbol, isPart)
		}
	} 
	
}
