package day_5

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
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

func Day_5(scanner *bufio.Scanner) (int64, int64) {
	startTime := time.Now()
	LoadData(scanner)
	dataLoadedTime := time.Now()
	fmt.Printf("Time to load data: %v s\n", (time.Since(startTime).Seconds()))
	// fmt.Println("Seeds loaded: ",len(seeds))
	// for i, m := range AllMaps{
	// 	fmt.Printf("Map %v: %v\n", i,(*m)[0])
	// }

	var location int64 = math.MaxInt64

	for i:=0; i < len(seeds) ; i +=2{
		seedStart := seeds[i]
		seedLength:= seeds[i+1]
		for j := int64(0); j < seedLength; j++{
			location = min(location, TravelSeedToLocation(seedStart+j))
		}
		
	}
	fmt.Printf("Time to find location: %v s\n", (time.Since(dataLoadedTime).Seconds()))
	return location, 0
}

func SelectMap(line string) *[]route{
	switch line{
	case "seed-to-soil map:":
		return &seedToSoilMap
	case "soil-to-fertilizer map:":
		return &soilToFertilizerMap
	case "fertilizer-to-water map:":
		return &fertilizerToWaterMap
	case "water-to-light map:":
		return &waterToLightMap
	case "light-to-temperature map:":
		return &lightToTemperatureMap
	case "temperature-to-humidity map:":
		return &temperatureToHumidityMap
	case "humidity-to-location map:":
		return &humidityToLocationMap
	default:
		return nil
	}
}

func ParseMapLine(line string) route{
	intSlice := []int64{}
	for _, vString := range strings.Split(line, " "){
		if vString == ""{
			continue
		}

		vInt, err:= strconv.ParseInt(vString,10,64)
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
		length: intSlice[0] - intSlice[1],
	}
}

func LoadData(scanner *bufio.Scanner) {
	
	scanner.Scan()
	firstLine := scanner.Text()
	firstLineSplit := strings.Split(firstLine, ":")
	for _, vString := range strings.Split(firstLineSplit[1], " "){
		vInt, err:= strconv.ParseInt(vString,10,64)
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

		// fmt.Println(line)
		*activeMap = append(*activeMap, ParseMapLine(line))
	}
}

func IsInRoute(value int64, r route) bool {
	b :=r.sourceStart <= value && value < r.sourceEnd
	// fmt.Printf("IsInRoute | value: %v, start: %v, end: %v, return: %v\n", value, r.sourceStart, r.sourceEnd, b)
	return b
}

func TravelSingle(value int64, m []route) (int64, bool){
	// fmt.Printf("moving value [%v] though map %v\n", value, m)
	for _, r := range m{
		if IsInRoute(value, r){
			return value + r.length, true
		}
	}
	return value, false
}

func TravelSeedToLocation(value int64) int64 {
	// exists := false
	for i := 0 ; i <= 6; i++{
		value, _ = TravelSingle(value, *AllMaps[i])
		// if !exists{
		// 	fmt.Println("pain exists")
		// 	break
		// }
	}
	return value
}