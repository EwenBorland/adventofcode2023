package day_13

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func ParseSolution() string {
	f, err := os.Open("day_13/input.txt")
	if err != nil {
		fmt.Println("file broke")
		return "error"
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	partOneAnswer, partTwoAnswer := Day_13(scanner)

	return fmt.Sprintf("Part 1 Answer: %v\nPart 2 Answer: %v\n", partOneAnswer, partTwoAnswer)
}

func Day_13(scanner *bufio.Scanner) (int64, int64) {
	patterns := [][][]rune{}
	newPattern := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)

		if line == "" {
			// fmt.Println("end of pattern detected")
			patterns = append(patterns, newPattern)
			newPattern = [][]rune{}
			continue
		}

		newLine := []rune{}
		for _, r := range line {
			newLine = append(newLine, r)
		}

		newPattern = append(newPattern, newLine)

	}

	if len(newPattern) != 0 {
		patterns = append(patterns, newPattern)
	}

	// fmt.Println("end of input")
	// fmt.Printf("%v patterns loaded\n", len(patterns))
	// for c, p := range patterns {
	// 	fmt.Printf("Pattern %v\n", c)
	// 	fmt.Println(p)
	// 	// for _, line := range p {
	// 	// 	fmt.Println(line)
	// 	// }
	// }

	summary := 0
	for _, pattern := range patterns {
		rowIndex, isMirror := FindRowMirror(pattern)
		if isMirror {
			summary += 100 * (rowIndex+1)
		}

		colIndex, isMirror := FindColumnMirror(pattern)
		if isMirror {
			summary += colIndex+1
		}
	}

	return int64(summary), 0
}

func AreColumnsEqual(indexA, indexB int, pattern [][]rune) bool {
	columnA := []rune{}
	columnB := []rune{}

	for _, row := range pattern {
		columnA = append(columnA, row[indexA])
		columnB = append(columnB, row[indexB])
	}
	// fmt.Printf("Comparing Columns: \n'%v'\n'%v'\n",columnA,columnB)
	// result := slices.Equal(columnA, columnB)
	// fmt.Println(result)
	return slices.Equal(columnA, columnB)
}

func AreRowsEqual(indexA, indexB int, pattern [][]rune) bool {
	// fmt.Printf("Comparing rows: \n'%v'\n'%v'\n",pattern[indexA],pattern[indexB])
	// result := slices.Equal(pattern[indexA], pattern[indexB])
	// fmt.Println(result)
	return slices.Equal(pattern[indexA], pattern[indexB])
}

func FindRowMirror(pattern [][]rune) (int, bool) {
	rowIndex := 0
	patternLen := len(pattern)
	isMirror := false
	for rowIndex = 0; rowIndex < patternLen-1; rowIndex++ {
		isMirror = true
		for compareDepth := patternLen; compareDepth >= 0; compareDepth-- {
			if rowIndex-compareDepth < 0 || rowIndex+compareDepth+2 > patternLen {
				continue
			}
			isMirror = AreRowsEqual(rowIndex-compareDepth, rowIndex+compareDepth+1, pattern)
			if !isMirror {
				break
			}
		}

		if isMirror {
			break
		}
	}

	return rowIndex, isMirror
}

func FindColumnMirror(pattern [][]rune) (int, bool) {
	colIndex := 0
	patternLen := len(pattern[0])
	isMirror := false
	for colIndex = 0; colIndex < patternLen-1; colIndex++ {
		isMirror = true
		for compareDepth := patternLen; compareDepth >= 0; compareDepth-- {
			// fmt.Printf("Column index: %v, Compare depth: %v\n", colIndex, compareDepth)
			if colIndex-compareDepth < 0 || colIndex+compareDepth+2 > patternLen {
				// fmt.Println("compare depth reached")
				continue
			}
			isMirror = AreColumnsEqual(colIndex-compareDepth, colIndex+compareDepth+1, pattern)
			if !isMirror {
				break
			}
		}

		if isMirror {
			break
		}
	}

	// fmt.Printf("Result of ColumnMirror | index: %v, mirror:%v\n", colIndex, isMirror)

	return colIndex, isMirror
}
