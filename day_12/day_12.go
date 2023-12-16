package day_12

import (
	"bufio"
	"fmt"
	"os"
)

func ParseSolution() string {
	f, err := os.Open("day_12/input.txt")
	if err != nil {
		fmt.Println("file broke")
		return "error"
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	partOneAnswer, partTwoAnswer := Day_12(scanner)

	return fmt.Sprintf("Part 1 Answer: %v\nPart 2 Answer: %v\n", partOneAnswer, partTwoAnswer)
}

func Day_12(scanner *bufio.Scanner) (int, int) {
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	return 0, 0
}

func CountArrangements(input string) int {
	// split := strings.Split(input, " ")
	// pattern := split[0]
	// springs := helpers.StringToInts(split[1], ",")
	return 0
}
