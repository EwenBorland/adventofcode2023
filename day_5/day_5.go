package day_5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseSolution() string {
	f, err := os.Open("day_5/input.txt")
	if err != nil {
		fmt.Println("file broke")
		return "error"
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	partOneAnswer, partTwoAnswer := Day_5(scanner)

	return fmt.Sprintf("Part 1 Answer: %v\nPart 2 Answer: %v\n", partOneAnswer, partTwoAnswer)
}

func Day_5(scanner *bufio.Scanner) (int, int) {
	LoadData(scanner)
	return 0, 0
}

func SelectMap(line string) *[]route{
	switch line{
	case "seed-to-soil map:":
		return &seedToSoilMap
	case "soil-to-fertilizer map:":
		return &seedToSoilMap
	case "fertilizer-to-water map:":
		return &seedToSoilMap
	case "water-to-light map:":
		return &seedToSoilMap
	case "light-to-temperature map:":
		return &seedToSoilMap
	case "temperature-to-humidity map:":
		return &seedToSoilMap
	case "humidity-to-location map:":
		return &seedToSoilMap
	default:
		return nil
	}
}

func ParseMapLine(line string) route{
	intSlice := []int{}
	for _, vString := range strings.Split(line, " "){
		if vString == ""{
			continue
		}

		vInt, err:= strconv.Atoi(vString)
		if err != nil{
			fmt.Println("pain")
		}

		intSlice = append(intSlice, vInt)
	}

	if len(intSlice) != 3{
		fmt.Println("pain: 2")
	}

	return route{
		destinationStart: intSlice[0],
		sourceStart: intSlice[1],
		destinationEnd: intSlice[0] + intSlice[2],
		sourceEnd: intSlice[1] + intSlice[2],
	}
}

func LoadData(scanner *bufio.Scanner) {
	
	scanner.Scan()
	firstLine := scanner.Text()
	firstLineSplit := strings.Split(firstLine, ":")
	for _, vString := range strings.Split(firstLineSplit[1], " "){
		vInt, err:= strconv.Atoi(vString)
		if err != nil{
			continue
		}
		seeds = append(seeds, vInt)
	}

	selectMap := false
	activeMap := &seedToSoilMap
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0{
			selectMap = true
			continue
		}

		if selectMap{
			activeMap = SelectMap(line)
			selectMap = false
			continue
		}

		fmt.Println(line)
		*activeMap = append(*activeMap, ParseMapLine(line))
	}
}