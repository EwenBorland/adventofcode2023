package main

import (
	"aoc2023/day_1"
	"aoc2023/day_2"
	"aoc2023/day_3"
	"aoc2023/day_4"
	"aoc2023/day_5"
	"fmt"
	"os"
	"strconv"
)

var activeDay int = 4

var dayFunctions map[int]func() string = map[int]func() string{
	1: day_1.ParseSolution,
	2: day_2.ParseSolution,
	3: day_3.ParseSolution,
	4: day_4.ParseSolution,
	5: day_5.ParseSolution,
}

func main() {
	if len(os.Args) > 1{
		argDay, err := strconv.Atoi(os.Args[1])
		if err != nil{
			fmt.Println("Argument not recognised")
		} else {
			activeDay = argDay
		}
	}
	fmt.Printf("Running solution for day %v\n", activeDay)
	result := dayFunctions[activeDay]()
	fmt.Println(result)
}