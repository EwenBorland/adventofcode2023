package day_10

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func ParseSolution() string {
	f, err := os.Open("day_10/input.txt")
	if err != nil {
		fmt.Println("file broke")
		return "error"
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	partOneAnswer, partTwoAnswer := Day_10(scanner)

	return fmt.Sprintf("Part 1 Answer: %v\nPart 2 Answer: %v\n", partOneAnswer, partTwoAnswer)
}

func Day_10(scanner *bufio.Scanner) (int, int) {
	groundMap := [][]rune{}
	lineIndex := 0
	var startI, startJ int = -1, -1

	for scanner.Scan() {
		line := scanner.Text()
		lineRunes, isStart := ParseLine(line)
		if isStart >= 0 {
			startI = lineIndex
			startJ = isStart
		}
		groundMap = append(groundMap, lineRunes)
		lineIndex++
	}

	if startI < 0 || startJ < 0 {
		fmt.Println("no starting indexes found")
	}

	horizLength := len(groundMap[0])

	pipeMap := [][]bool{}
	for i:= 0; i < len(groundMap); i++{
		pipeMap = append(pipeMap, make([]bool, horizLength))

	}

	fmt.Printf("starting at: {%v,%v}\n", startI, startJ)

	ci, cj := StartStep(startI, startJ, groundMap)
	// fmt.Printf("Moving from {%v,%v} to {%v,%v}\n", startI, startJ, ci, cj)
	pi, pj := startI, startJ
	newi, newj := 0, 0
	totalSteps := 1

	maxSteps := len(groundMap) * len(groundMap[0])
	for totalSteps <= maxSteps {
		totalSteps += 1
		newi, newj = Step(ci, cj, pi, pj, groundMap, pipeMap)
		// fmt.Printf("Moving from {%v,%v} to {%v,%v}\n", ci, cj, newi, newj)
		if newi < 0 {
			break
		}
		pi, pj = ci, cj
		ci, cj = newi, newj
	}

	fmt.Println(totalSteps, "steps to return to start")
	fmt.Println(maxSteps, "step limit")
	
	enclosedCount := 0
	for i,v := range groundMap{
		enclosedCount += CountEnclosed(v, pipeMap[i])
	}

	return totalSteps / 2, enclosedCount
}

func ParseLine(line string) ([]rune, int) {
	lineRunes := []rune{}
	start := -1
	for j, v := range line {
		lineRunes = append(lineRunes, v)
		if v == 'S' {
			start = j
		}
	}
	return lineRunes, start
}

func Step(ci, cj, pi, pj int, groundMap [][]rune, pipeMap [][]bool) (int, int) {
	currentSymbol := groundMap[ci][cj]
	pipeMap[ci][cj] = true	

	switch currentSymbol {
	case '|':
		return ci + (ci - pi), cj
	case '-':
		return ci, cj + (cj - pj)
	case 'L':
		return ci + (ci - pi) - 1, cj + (cj - pj) + 1
	case 'J':
		return ci + (ci - pi) - 1, cj + (cj - pj) - 1
	case '7':
		return ci + (ci - pi) + 1, cj + (cj - pj) - 1
	case 'F':
		return ci + (ci - pi) + 1, cj + (cj - pj) + 1
	case 'S':
		return -1, -1
	default:
		fmt.Println("invalid symbol:",currentSymbol)
		return -1, -1
	}
}

func StartStep(ci, cj int, groundMap [][]rune) (int, int) {
	if slices.Contains([]rune{'|', '7', 'F'}, groundMap[ci+1][cj]) {
		//above
		return ci + 1, cj
	} else if slices.Contains([]rune{'|', 'L', 'J'}, groundMap[ci-1][cj]) {
		//below
		return ci - 1, cj
	} else if slices.Contains([]rune{'-', 'J', '7'}, groundMap[ci][cj+1]) {
		//right
		return ci, cj + 1
	} else {
		//left
		return ci, cj - 1
	}
}

func CountEnclosed(groundRow []rune,pipeRow []bool) int{
	enclosed := false
	enclosedCount := 0
	previousSymbol := '.'
	for i, isPipe := range pipeRow{
		if isPipe{
			if EnclosedToggle(groundRow[i],&previousSymbol){
				enclosed= !enclosed
			}
		} else {
			if enclosed{
				enclosedCount++
			}
		}
	}
	return enclosedCount
}

func EnclosedToggle(symbol rune, previousSymbol *rune) bool{
	switch symbol {
	case '|':
		return true
	case 'L':
		*previousSymbol = symbol
		return false
	case 'J':
		if *previousSymbol == 'F'{
			return true 
		}
		*previousSymbol = 'J'
		return false
	case '7':
		if *previousSymbol == 'L'{
			return true 
		}
		*previousSymbol = '7'
		return false
	case 'F':
		*previousSymbol = symbol
		return false
	case 'S':
		return true // I'm assuming the starting symbol is a vertical pipe
	case '-':
		return false
	default:
		fmt.Println("invalid symbol:",symbol)
		return false
	}
}
