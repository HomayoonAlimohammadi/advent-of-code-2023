package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed test.txt
var testDataB []byte

//go:embed input.txt
var inputDataB []byte

type Map struct {
	dstStart int
	srcStart int
	length   int
}

type SeedFinder struct {
	seeds               []int
	seedToSoils         []Map
	soilToFertilizers   []Map
	fertilizerToWaters  []Map
	waterToLights       []Map
	lightToTemps        []Map
	tempToHumidities    []Map
	humidityToLocations []Map
}

func (sf *SeedFinder) MapSeeds() []int {

	locs := make([]int, len(sf.seeds))

	for i, seed := range sf.seeds {
		loc := sf.HumidityToLocation(
			sf.TempToHumidity(
				sf.LightToTemp(
					sf.WaterToLight(
						sf.FertilizerToWater(
							sf.SoilToFertilizer(
								sf.SeedToSoil(seed),
							),
						),
					),
				),
			),
		)

		locs[i] = loc
	}

	return locs
}

func (sf *SeedFinder) SeedToSoil(src int) int {
	for _, m := range sf.seedToSoils {
		if m.srcStart <= src && src <= m.srcStart+m.length {
			diff := src - m.srcStart
			return m.dstStart + diff
		}
	}

	return src
}

func (sf *SeedFinder) SoilToFertilizer(src int) int {
	for _, m := range sf.soilToFertilizers {
		if m.srcStart <= src && src <= m.srcStart+m.length {
			diff := src - m.srcStart
			return m.dstStart + diff
		}
	}

	return src
}

func (sf *SeedFinder) FertilizerToWater(src int) int {
	for _, m := range sf.fertilizerToWaters {
		if m.srcStart <= src && src <= m.srcStart+m.length {
			diff := src - m.srcStart
			return m.dstStart + diff
		}
	}

	return src
}

func (sf *SeedFinder) WaterToLight(src int) int {
	for _, m := range sf.waterToLights {
		if m.srcStart <= src && src <= m.srcStart+m.length {
			diff := src - m.srcStart
			return m.dstStart + diff
		}
	}

	return src
}

func (sf *SeedFinder) LightToTemp(src int) int {
	for _, m := range sf.lightToTemps {
		if m.srcStart <= src && src <= m.srcStart+m.length {
			diff := src - m.srcStart
			return m.dstStart + diff
		}
	}

	return src
}

func (sf *SeedFinder) TempToHumidity(src int) int {
	for _, m := range sf.tempToHumidities {
		if m.srcStart <= src && src <= m.srcStart+m.length {
			diff := src - m.srcStart
			return m.dstStart + diff
		}
	}

	return src
}

func (sf *SeedFinder) HumidityToLocation(src int) int {
	for _, m := range sf.humidityToLocations {
		if m.srcStart <= src && src <= m.srcStart+m.length {
			diff := src - m.srcStart
			return m.dstStart + diff
		}
	}

	return src
}

func main() {
	input := string(inputDataB)

	seedFinder := SeedFinder{}
	currentMap := &seedFinder.seedToSoils // Start with the first map

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue // Skip empty lines
		}

		if line == "seed-to-soil map:" {
			currentMap = &seedFinder.seedToSoils
		} else if line == "soil-to-fertilizer map:" {
			currentMap = &seedFinder.soilToFertilizers
		} else if line == "fertilizer-to-water map:" {
			currentMap = &seedFinder.fertilizerToWaters
		} else if line == "water-to-light map:" {
			currentMap = &seedFinder.waterToLights
		} else if line == "light-to-temperature map:" {
			currentMap = &seedFinder.lightToTemps
		} else if line == "temperature-to-humidity map:" {
			currentMap = &seedFinder.tempToHumidities
		} else if line == "humidity-to-location map:" {
			currentMap = &seedFinder.humidityToLocations
		} else if strings.HasPrefix(line, "seeds:") {
			// Parse seed numbers
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				seedFinder.seeds = parseIntArray(parts[1])
			}
		} else {
			// Parse map entries
			entry := parseIntArray(line)
			if len(entry) == 3 {
				*currentMap = append(*currentMap, Map{
					dstStart: entry[0],
					srcStart: entry[1],
					length:   entry[2],
				})
			}
		}
	}

	locs := seedFinder.MapSeeds()

	fmt.Println(locs)

	minLoc := int(math.Inf(1))
	for _, l := range locs {
		minLoc = min(l, minLoc)
	}

	fmt.Println(minLoc)
}

func parseIntArray(line string) []int {
	parts := strings.Fields(line)
	var result []int
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err == nil {
			result = append(result, num)
		}
	}
	return result
}
