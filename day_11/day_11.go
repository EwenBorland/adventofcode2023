package day_11

import (
	"bufio"
	"fmt"
	"os"
)

func ParseSolution() string {
	f, err := os.Open("day_11/input.txt")
	if err != nil {
		fmt.Println("file broke")
		return "error"
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	partOneAnswer, partTwoAnswer := Day_11(scanner)

	return fmt.Sprintf("Part 1 Answer: %v\nPart 2 Answer: %v\n", partOneAnswer, partTwoAnswer)
}

func Day_11(scanner *bufio.Scanner) (int64, int64) {
	galaxies := []map[int]galaxy{}
	columnsHaveGalaxies := map[int]bool{}

	currentLine := int64(0)
	line := ""
	for scanner.Scan() {
		line = scanner.Text()

		currentGalaxies, hasGalaxies := ParseLine(line, currentLine, &columnsHaveGalaxies)

		galaxies = append(galaxies, currentGalaxies)
		if !hasGalaxies {
			// currentLine++
			currentLine += 999999
		}
		currentLine++
	}

	columns := len(line)
	counter := int64(0)
	// fmt.Println(columnsHaveGalaxies)
	// fmt.Println(galaxies)
	for i := 0; i < columns; i++ {
		if !columnsHaveGalaxies[i] {
			// counter++
			counter += 999999
			continue
		}

		for j, row := range galaxies {
			g, ok := row[i]
			if ok {
				// fmt.Println(g)
				g.horizontalIndex += counter
				// fmt.Println(g)
				galaxies[j][i] = g
			}
		}
	}
	// fmt.Println(galaxies)

	// expansionLine := "................."
	// lineN := 0
	// totalCols := columns+counter
	// for _, row := range galaxies{
	// 	newLine := ""
	// 	for j:=0 ; j < totalCols; j++{
	// 		g, ok := row[j]
	// 		if !ok{
	// 			newLine = newLine + "."
	// 			continue
	// 		}
	// 		if g.verticalIndex != lineN{
	// 			fmt.Println(expansionLine)
	// 			lineN++
	// 		}
	// 		newLine = newLine + "#"
	// 	}
	// 	fmt.Println(newLine)
	// 	lineN++
	// }

	newGalaxies := []galaxy{}
	for _, row := range galaxies {
		for _, g := range row {
			newGalaxies = append(newGalaxies, g)
		}
	}

	pairDistCount := int64(0)
	for i := 0; i < len(newGalaxies); i++ {
		currentG := newGalaxies[i]
		for _, g := range newGalaxies[i+1:] {
			pairDistCount += GalaxyDistance(currentG, g)
		}
	}
	return pairDistCount, 0
}

type galaxy struct {
	verticalIndex   int64
	horizontalIndex int64
}

func ParseLine(line string, lineNumber int64, columnsHaveGalaxies *map[int]bool) (map[int]galaxy, bool) {
	gs := map[int]galaxy{}
	hasGalaxies := false
	for i, v := range line {
		if v == '#' {
			hasGalaxies = true
			gs[i] = galaxy{verticalIndex: lineNumber, horizontalIndex: int64(i)}
			(*columnsHaveGalaxies)[i] = true
		}
	}
	return gs, hasGalaxies
}

func GalaxyDistance(g1, g2 galaxy) int64 {
	v := g1.verticalIndex - g2.verticalIndex
	if v < 0 {
		v = -v
	}
	h := g1.horizontalIndex - g2.horizontalIndex
	if h < 0 {
		h = -h
	}
	return v + h
}
