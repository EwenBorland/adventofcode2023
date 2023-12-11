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

	pipeMap, enclosedMap := [][]bool{}, [][]bool{}
	for i:= 0; i < len(groundMap); i++{
		pipeMap = append(pipeMap, make([]bool, horizLength))
		enclosedMap = append(enclosedMap, make([]bool, horizLength))
	}

	fmt.Printf("starting at: {%v,%v}\n", startI, startJ)

	ci, cj := StartStep(startI, startJ, groundMap)
	// fmt.Printf("Moving from {%v,%v} to {%v,%v}\n", startI, startJ, ci, cj)
	pi, pj := startI, startJ
	newi, newj := 0, 0
	totalSteps := 1
	isOutside := false
	maxSteps := len(groundMap) * len(groundMap[0])
	for totalSteps <= maxSteps {
		totalSteps += 1
		newi, newj = Step(ci, cj, pi, pj, groundMap, pipeMap, enclosedMap, &isOutside)
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
	pipeCount := 0
	for i, v := range groundMap{
		for j := range v{
			if pipeMap[i][j]{
				pipeCount++
				continue
			} 
			if enclosedMap[i][j] {
				enclosedCount++
			}
		}


	}
	fmt.Println("Pipes: ",pipeCount)
	fmt.Println("enclosed: ", enclosedCount)
	fmt.Println("area: ", maxSteps)

	if isOutside {
		enclosedCount = maxSteps - pipeCount - enclosedCount
		fmt.Println("enclosed 2 : ", enclosedCount)
	}
	//for part two, determine direction, store tiles to right of direction and tile that are in loop.
	// at end, remove duplicates and loop tiles from the 'right' tiles, the remaining 'righ' tiles are either inside the loop or outside
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

func Step(ci, cj, pi, pj int, groundMap [][]rune, pipeMap , enclosedMap [][]bool, outside *bool) (int, int) {
	currentSymbol := groundMap[ci][cj]
	pipeMap[ci][cj] = true

	enclosedi, enclosedj := ci+cj-pj, cj+pi-ci
	if enclosedi< 0|| enclosedj < 0 || enclosedi >= len(enclosedMap)|| enclosedj >= len(enclosedMap[0]){
		*outside = true
	} else{
		enclosedMap[ci+cj-pj][cj+pi-ci] = true
	}
	

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
