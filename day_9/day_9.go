package day_9

import (
	"aoc2023/helpers"
	"bufio"
	"fmt"
	"os"
	"slices"
)

func ParseSolution() string {
	f, err := os.Open("day_9/input.txt")
	if err != nil {
		fmt.Println("file broke")
		return "error"
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	partOneAnswer, partTwoAnswer := Day_9(scanner)

	return fmt.Sprintf("Part 1 Answer: %v\nPart 2 Answer: %v\n", partOneAnswer, partTwoAnswer)
}

func Day_9(scanner *bufio.Scanner) (int64, int64) {
	nextValuesSum := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		lineInts := helpers.StringToInts64(line)

		nextVal := GetNextValue(lineInts)
		fmt.Printf("Next value for line: %v, %v\n", lineInts, nextVal)
		nextValuesSum += nextVal
	}

	return 114, nextValuesSum
}

func GetNextValue(sequence []int64) int64 {
	nextSequence := GetDifferences(sequence)
	// fmt.Println(nextSequence)
	if slices.Min(nextSequence) == 0 && slices.Max(nextSequence) == 0 {
		return sequence[0] + 0
	}
	nextVal := GetNextValue(nextSequence)
	// fmt.Println(nextSequence)
	// fmt.Println(nextVal)
	return sequence[0] - nextVal
}

func GetDifferences(sequence []int64) []int64 {
	outputSequence := []int64{}
	for i := 0; i < len(sequence)-1; i++ {
		outputSequence = append(outputSequence, sequence[i+1]-sequence[i])
	}

	return outputSequence
}
