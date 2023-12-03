package main

import (
	"aoc2023/day_1"
	"aoc2023/day_2"
	"aoc2023/day_3"
	"fmt"
)

const currentDay int = 2

var dayFunctions map[int]func() string = map[int]func() string{
	1: day_1.ParseSolution,
	2: day_2.ParseSolution,
	3: day_3.ParseSolution,
}

func main() {
	fmt.Printf("Running solution for day %v\n", currentDay)
	result := dayFunctions[currentDay]()
	fmt.Println(result)
}