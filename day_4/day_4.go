package day_4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ParseSolution() string {
	f, err := os.Open("day_4/input.txt")
	if err != nil {
		fmt.Println("file broke")
		return "error"
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	partOneAnswer, partTwoAnswer := Day_4(scanner)

	return fmt.Sprintf("Part 1 Answer: %v\nPart 2 Answer: %v\n", partOneAnswer, partTwoAnswer)
}

func Day_4(scanner *bufio.Scanner) (int, int) {
	score := 0
	cards := 0
	instanceNums := map[int]int{}
	lineNum := -1

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		// fmt.Println(line)
		winNums, youNums, err := ParseLine(line)
		if err != nil{
			log.Fatal(err)
		}
		winners := 0
		for _, y := range youNums{
			if slices.Contains(winNums, y){
				// fmt.Printf("%v is in %v\n",y, winNums)
				winners++
				continue
			}
			// fmt.Printf("%v is NOT in %v\n",y, winNums)
		}

		lineScore := 0
		if winners > 0{
			lineScore = IntPow(2,winners-1)
			// fmt.Printf("winners: %v, worth %v\n",winners, lineScore)
			winnersSlice := make([]int, winners)
			// fmt.Printf("winnersSlice len %v\n",len(winnersSlice))
			for i := range winnersSlice{
				// fmt.Printf("Adding %v copies of card %v\n",1*(instanceNums[lineNum]+1),lineNum+1+i)
				instanceNums[lineNum+1+i] += 1*(instanceNums[lineNum]+1)
			}
		}
		// fmt.Printf("scoring %v instances of card %v\n", instanceNums[lineNum]+1, lineNum)
		score += lineScore*(instanceNums[lineNum]+1)
		cards += instanceNums[lineNum]+1

	}

	return score, cards

}

func ParseLine(line string) ([]int, []int, error){
	// line =Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	lineSplit := strings.Split(line, ":")
	// lineSplit =["Card 1", "41 48 83 86 17 | 83 86  6 31 17  9 48 53"]
	lineSplit = strings.Split(lineSplit[1], "|")
	// lineSplit =["41 48 83 86 17 "," 83 86  6 31 17  9 48 53"]

	winNumStrings := strings.Split(lineSplit[0]," ")
	youNumString := strings.Split(lineSplit[1]," ")

	winNums, youNums := []int{}, []int{}

	for _, s := range winNumStrings{
		// fmt.Println(s)
		i, err := strconv.Atoi(s)
		if err != nil{
			// fmt.Printf("Error parsing '%v' as int, err: %v\n", s, err)
			continue
		}
		winNums = append(winNums, i)
	}

	for _, s := range youNumString{
		// fmt.Println(s)
		i, err := strconv.Atoi(s)
		if err != nil{
			// fmt.Printf("Error parsing '%v' as int, err: %v\n", s, err)
			continue
		}
		youNums = append(youNums, i)
	}

	return winNums, youNums, nil
}

func IntPow(n, m int) int {
    if m == 0 {
        return 1
    }
    result := n
    for i := 2; i <= m; i++ {
        result *= n
    }
    return result
}
