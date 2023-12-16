package day_14

import (
	"bufio"
	"fmt"
	"os"
)

func ParseSolution() string {
	f, err := os.Open("day_14/input.txt")
	if err != nil {
		fmt.Println("file broke")
		return "error"
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	partOneAnswer, partTwoAnswer := Day_14(scanner)

	return fmt.Sprintf("Part 1 Answer: %v\nPart 2 Answer: %v\n", partOneAnswer, partTwoAnswer)
}

func Day_14(scanner *bufio.Scanner) (int, int) {

	dish := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		newLine := []rune{}
		for _, r := range line {
			newLine = append(newLine, r)
		}

		dish = append(dish, newLine)

	}
	// fmt.Println("Dish loaded:")
	// PrintDish(dish)

	// fmt.Println("Dish tilted north:")
	northDish := TiltNorth(dish)
	// PrintDish(northDish)

	load := MeasureLoad(northDish)

	return load, 0
}

// for part 2, will need TiltEast/West/South funcs
// when 'rolling', after every tilt north, Flatten the dish to a string and use that as a key in a map
// then ok check to see if we have seen that key before, this marks the start/end of a loop
// then some lowest common multiple shenanigans to figure out which iteration in the loop will correspond to the billionth overall iteration

func PrintDish(dish [][]rune){
	fmt.Println("------------")
	for _,row := range dish{
		line := ""
		for _, r :=range row{
			line = line + string(r)
		}
		fmt.Println(line)
	}
	fmt.Println("------------")
}

func TiltNorth(dish [][]rune) [][]rune {
	
	for rowIndex, row := range dish {
		for colIndex, r := range row {
			if r != 'O' {
				continue
			}
			newRowIndex := rowIndex
			for rollrowIndex := rowIndex - 1; rollrowIndex >= 0; rollrowIndex-- {
				if dish[rollrowIndex][colIndex] == '#' || dish[rollrowIndex][colIndex] == 'O' {
					newRowIndex = rollrowIndex + 1
					break
				} else if rollrowIndex == 0{
					newRowIndex = 0
				}
			}
			dish[rowIndex][colIndex] = '.'
			dish[newRowIndex][colIndex] = 'O'
		}
	}

	return dish
}

func MeasureLoad(dish [][]rune) int{
	totalRows := len(dish)
	load := 0

	for rowIndex, row := range dish{
		for _, r := range row{
			if r != 'O'{
				continue
			}
			load += totalRows - rowIndex
		}
	}
	return load
}
