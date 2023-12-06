package day_3

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func ParseSolution() string {
	f, err := os.Open("day_3/input.txt")
	if err != nil {
		fmt.Println("file broke")
		return "error"
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	validCount, gearScore := Day_3(scanner)

	return fmt.Sprintf("Valid Sum: %v\nGear Ratio: %v", validCount, gearScore)
}

/*

idea 1:
masking maybe

initialise the 0th and 1st line with falses

if a symbol is parsed then set the appropriate indexes in the mask true

apply the mask to the parts somehow

idea 2:
the symbols only affect parts that are on adjacent lines, parse line n for parts, then parse lines n-1,n,and n+1 for symbols and match indices

idea 3:
for each line, if there's a number, store it in a list with the line num and all indices. if there's a symbol calculate the suttounding indices and add them to a map of line: valid indices

*/

type Part struct {
	Line    int
	Value   int
	Indices []int
}

type Symbol struct {
	Line  int
	Index int
	Value rune
}

var intsAsRunes []rune = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

// IsSymbolOrPart returns (c==Symbol, c==Part)
func IsSymbolOrPart(c rune) (bool, bool) {

	if c == '.' || c == ' ' {
		// fmt.Printf("IsSymbolOrPart| '%c' identified as neither\n", c)
		return false, false
	}
	if slices.Contains(intsAsRunes, c) {
		// fmt.Printf("IsSymbolOrPart| '%c' identified as Part\n", c)
		return false, true
	}
	// fmt.Printf("IsSymbolOrPart| '%c' identified as Symbol\n", c)
	return true, false
}

func ParseLine(line string, lineNumber int, parts *[]Part, symbols *[]Symbol, validIndices *map[int][]int) {
	parsingPart := false
	fmt.Printf("Parsing line {%v}: '%v'\n", lineNumber, line)
	for i, char := range line {
		isSymbol, isPart := IsSymbolOrPart(char)
		if isSymbol {
			parsingPart = false
			ParseSymbol(i, lineNumber, char, symbols, validIndices)
		} else if isPart {
			if parsingPart {
				continue
			}
			parsingPart = true
			ParsePart(i, lineNumber, line, parts)
		} else {
			parsingPart = false
		}

	}
}

func ParseSymbol(x, y int, char rune, symbols *[]Symbol, validIndices *map[int][]int) {
	s := Symbol{
		Line:    y,
		Value: char,
		Index: x,
	}
	*symbols = append(*symbols, s)

	minY := max(0, y-1)
	maxY := y + 1
	minX := max(0, x-1)
	maxX := x + 1
	for ln := minY; ln <= maxY; ln++ {
		for index := minX; index <= maxX; index++ {
			if !slices.Contains((*validIndices)[ln], index) {
				(*validIndices)[ln] = append((*validIndices)[ln], index)
			}
		}

	}
}

func ParsePart(x, y int, line string, parts *[]Part) {
	p := Part{
		Line:    y,
		Value:   0,
		Indices: []int{x},
	}
	x++
	for _, char := range line[x:] {
		_, isPart := IsSymbolOrPart(char)
		if !isPart {
			break
		}
		p.Indices = append(p.Indices, x)
		x++
	}
	p.Value, _ = strconv.Atoi(line[p.Indices[0] : p.Indices[len(p.Indices)-1]+1])
	*parts = append(*parts, p)
}

func IsPartValid(p Part, v map[int][]int) bool {
	// fmt.Printf("Is part indices:[%v] in map indices:[%v]? ", p.Indices, v[p.Line])
	for _, ind := range p.Indices {
		if slices.Contains(v[p.Line], ind) {
			return true
		}
	}
	return false
}

func Day_3(scanner *bufio.Scanner) (int, int) {
	lineN := 0
	parts := []Part{}
	symbols := []Symbol{}
	validMap := map[int][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		ParseLine(line, lineN, &parts, &symbols, &validMap)

		lineN++
	}
	fmt.Printf("Found %v parts\n", len(parts))
	fmt.Printf("Found %v symbols\n", len(symbols))
	partSum := 0
	for _, p := range parts {
		if IsPartValid(p, validMap) {
			partSum += p.Value
		}
	}

	// sort parts into lines
	var partLineMap map[int][]Part = map[int][]Part{}
	for _, p := range parts{
		partLineMap[p.Line] = append(partLineMap[p.Line], p)
	}

	gearScore := 0

	for _, s := range symbols {
		adjacents := 0
		adjacentParts := []Part{}
		if s.Value != '*'{
			continue
		}
		sliceToCheck := append(partLineMap[s.Line],partLineMap[s.Line+1]...)
		if s.Line > 0{
			sliceToCheck =append(sliceToCheck,partLineMap[max(0, s.Line-1)]...)
		}
		
		for _, p := range sliceToCheck{
			if slices.Contains(p.Indices,s.Index) || slices.Contains(p.Indices,s.Index-1) || slices.Contains(p.Indices,s.Index+1){
				adjacents+=1
				adjacentParts = append(adjacentParts, p)
			}
		}

		if adjacents != 2{
			continue
		}

		gearScore += adjacentParts[0].Value * adjacentParts[1].Value

	}

	return partSum, gearScore

}
