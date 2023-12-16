package day_6

import (
	"aoc2023/helpers"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseSolution() string {
	f, err := os.Open("day_6/input.txt")
	if err != nil {
		fmt.Println("file broke")
		return "error"
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	partOneAnswer, partTwoAnswer := Day_6(scanner)

	return fmt.Sprintf("Part 1 Answer: %v\nPart 2 Answer: %v\n", partOneAnswer, partTwoAnswer)
}

func Day_6(scanner *bufio.Scanner) (int, int) {
	scanner.Scan()
	lineSplit := strings.Split(scanner.Text(), ":")
	timesP1 := helpers.StringToInts(lineSplit[1], " ")
	timeP2, err := strconv.ParseInt(strings.ReplaceAll(lineSplit[1], " ", ""),10,64)
	if err != nil{
		fmt.Println(err)
	}

	scanner.Scan()
	lineSplit = strings.Split(scanner.Text(), ":")
	distancesP1 := helpers.StringToInts(lineSplit[1], " ")
	distanceP2, err := strconv.ParseInt(strings.ReplaceAll(lineSplit[1], " ", ""),10,64)
	if err != nil{
		fmt.Println(err)
	}

	options := []int{}
	for i := 0; i < len(timesP1); i++ {
		a := EvaluateRace(timesP1[i], distancesP1[i])
		options = append(options, a)
	}

	part1Answer := 1
	for _, v := range options {
		part1Answer *= v
	}

	part2Answer := EvaluateRace64(timeP2, distanceP2)

	return part1Answer, part2Answer
}

func EvaluateRace(time int, distance int) int {
	minSpeed := 1
	winning := false
	for !winning {
		minSpeed++
		winning = (time-minSpeed)*minSpeed > distance
	}

	fmt.Printf("Min speed for time:%v dist:%v is speed:%v\n", time, distance, minSpeed)

	options := 0
	for speed := minSpeed; (time-speed)*speed > distance; speed++ {
		options++
	}
	return options
}

func EvaluateRace64(time int64, distance int64) int {
	minSpeed := int64(1)
	winning := false
	for !winning {
		minSpeed++
		winning = (time-minSpeed)*minSpeed > distance
	}

	fmt.Printf("Min speed for time:%v dist:%v is speed:%v\n", time, distance, minSpeed)

	options := 0
	for speed := minSpeed; (time-speed)*speed > distance; speed++ {
		options++
	}
	return options
}
