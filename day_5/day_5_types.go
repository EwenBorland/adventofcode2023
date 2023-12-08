package day_5

import "fmt"

var (
	seeds                    = []int64{}
	seedToSoilMap            = []route{}
	soilToFertilizerMap      = []route{}
	fertilizerToWaterMap     = []route{}
	waterToLightMap          = []route{}
	lightToTemperatureMap    = []route{}
	temperatureToHumidityMap = []route{}
	humidityToLocationMap    = []route{}

	AllMaps map[int]*[]route = map[int]*[]route{
		0: &seedToSoilMap,
		1: &soilToFertilizerMap,
		2: &fertilizerToWaterMap,
		3: &waterToLightMap,
		4: &lightToTemperatureMap,
		5: &temperatureToHumidityMap,
		6: &humidityToLocationMap,
	}
)

const ()

type route struct {
	sourceStart      int64
	destinationStart int64
	sourceEnd        int64
	destinationEnd   int64
	length           int64
}

func (r route) String() string {
	return fmt.Sprintf("{sourceStart: %v, destStart: %v, sourceEnd: %v, destEnd: %v,len: %v}",r.sourceStart,r.destinationStart,r.sourceEnd,r.destinationEnd,r.length)
}
