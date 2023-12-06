package day_5

var (
	seeds = []int{}
	seedToSoilMap = []route{}
	soilToFertilizerMap = []route{}
	fertilizerToWaterMap = []route{}
	waterToLightMap = []route{}
	lightToTemperatureMap = []route{}
	temperatureToHumidityMap = []route{}
	humidityToLocationMap = []route{}
)

const (

)

type route struct{
	sourceStart int
	destinationStart int
	sourceEnd int
	destinationEnd int
}

