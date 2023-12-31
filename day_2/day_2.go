package day_2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseSolution() string {
	validCount, powerCount  := day_2()
	return fmt.Sprintf("Valid Sum: %v\nPower Sum: %v\n", validCount, powerCount)
}

func day_2() (int, int) {
	f, err := os.Open("day_2/input.txt")
	if err != nil{
		fmt.Println("file broke")
		return 0, 0
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	gameID := 1
	validCount := 0
	powerCount := 0

	limitsColour := map[string]int{
		"red": 12,
		"green": 13,
		"blue": 14,
	}
	
	for scanner.Scan() {
		maxColour := map[string]int{}

		line := scanner.Text()
		// fmt.Println(line)
		lineSplit := strings.Split(line, ":")
		for _, set := range strings.Split(lineSplit[1], ";"){
			for _, colourSplit := range strings.Split(set, ","){
				colNameVal := strings.Split(colourSplit, " ")

				val, err := strconv.Atoi(colNameVal[1])
				if err != nil{
					log.Fatal(err)
				}

				// fmt.Printf("Col:%v, Val: %v, maxCol: %v\n", colNameVal[2],colNameVal[1], maxColour[colNameVal[2]])
				if maxColour[colNameVal[2]] < val{
					maxColour[colNameVal[2]] = val
				}
			}
		}

		

		valid := true
		for colourName, limit := range limitsColour{
			if maxColour[colourName] > limit{
				// fmt.Printf("Invalid due to %v\n", colourName)
				valid = false
				break
			}
		}

		if valid{
			validCount += gameID
		}

		powerCount += maxColour["red"] * maxColour["blue"] * maxColour["green"]

		gameID++
	}

	return validCount, powerCount

}
