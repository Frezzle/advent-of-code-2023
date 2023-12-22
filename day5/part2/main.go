package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	parts := strings.Split(input, "\n\n")
	seedRanges := getNums(parts[0][7:])
	seedToSoil := newMap(parts[1][18:])
	soilToFertilizer := newMap(parts[2][24:])
	fertilizerToWater := newMap(parts[3][25:])
	waterToLight := newMap(parts[4][20:])
	lightToTemperature := newMap(parts[5][26:])
	temperatureToHumidity := newMap(parts[6][29:])
	humidityToLocation := newMap(parts[7][26:])

	closestLocation := math.MaxInt
	for i := 0; i < len(seedRanges); i += 2 {
		firstSeed := seedRanges[i]
		seedsRange := seedRanges[i+1]
		for seed := firstSeed; seed < firstSeed+seedsRange; seed++ {
			soil := seedToSoil.GetValue(seed)
			fert := soilToFertilizer.GetValue(soil)
			water := fertilizerToWater.GetValue(fert)
			light := waterToLight.GetValue(water)
			temp := lightToTemperature.GetValue(light)
			humi := temperatureToHumidity.GetValue(temp)
			loc := humidityToLocation.GetValue(humi)
			if loc < closestLocation {
				closestLocation = loc
			}
		}
	}

	fmt.Println("Closest location:", closestLocation)
}

func newMap(s string) *Map {
	m := new(Map)
	rangeStrings := strings.Split(s, "\n")
	for _, rangeString := range rangeStrings {
		nums := getNums(rangeString)
		m.Ranges = append(m.Ranges, Range{
			Source:      nums[1],
			Destination: nums[0],
			Length:      nums[2],
		})
	}
	return m
}

func getNums(s string) []int {
	numStrings := strings.Split(s, " ")
	nums := make([]int, 0, len(numStrings))
	for _, x := range numStrings {
		num, _ := strconv.Atoi(x)
		nums = append(nums, num)
	}
	return nums
}

type Map struct {
	Ranges []Range
}

type Range struct {
	Source      int
	Destination int
	Length      int
}

func (m *Map) GetValue(key int) int {
	for _, r := range m.Ranges {
		if key >= r.Source && key < r.Source+r.Length {
			return key + r.Destination - r.Source
		}
	}
	return key
}
