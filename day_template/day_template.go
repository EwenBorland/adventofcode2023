package day_template

import (
	"bufio"
	"fmt"
	"os"
)

func ParseSolution() string {
	f, err := os.Open("day_template/input.txt")
	if err != nil {
		fmt.Println("file broke")
		return "error"
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	partOneAnswer, partTwoAnswer := Day_template(scanner)

	return fmt.Sprintf("Part 1 Answer: %v\nPart 2 Answer: %v\n", partOneAnswer, partTwoAnswer)
}

func Day_template(scanner *bufio.Scanner) (int, int) {
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	return 0, 0
}
