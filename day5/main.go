package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

// part1: 836040384
// part2: 358218184 too high
// part2: 10834440 but very slow algorithm!
func main() {
	//entries := getEntries("part1-example.txt")
	entries := getEntries("part1-input.txt")

	location := getNearestLocation2(entries)
	fmt.Println("Nearest location: ", location)
}

func getNearestLocation2(entries []string) int {
	seeds := getSeeds(entries[0])
	mappers := getMappers(entries[2:])

	nearest := math.MaxInt
	for i := 0; i+1 < len(seeds); i += 2 {
		for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
			if location := getLocation(mappers, seed); location < nearest {
				nearest = location
			}
		}
	}

	return nearest
}

// part1: 836040384
func getNearestLocation(entries []string) int {
	seeds := getSeeds(entries[0])
	mappers := getMappers(entries[2:])

	var locations []int
	for _, seed := range seeds {
		locations = append(locations, getLocation(mappers, seed))
	}
	slices.Sort(locations)

	return locations[0]
}

type mapping struct {
	dst         int
	src         int
	rangeLength int
}

type mapper []mapping

func (m mapper) dst(v int) int {
	for _, ma := range m {
		if v >= ma.src && v < ma.src+ma.rangeLength {
			return v - ma.src + ma.dst
		}
	}

	return v
}

func getLocation(mappers []mapper, seed int) int {
	val := seed

	for _, mapper := range mappers {
		val = mapper.dst(val)
	}

	return val
}

func getMappers(entries []string) []mapper {
	var mappers []mapper
	var currentMapper mapper

	for i, entry := range entries {
		// New mapper
		if strings.HasSuffix(entry, " map:") {
			currentMapper = mapper{}
			continue
		}

		// Mappers separator
		if entry == "" {
			mappers = append(mappers, currentMapper)
			continue
		}

		// Else we are in a mapping line
		mappingValues := strings.Split(entry, " ")
		currentMapper = append(currentMapper, mapping{
			dst:         toInt(mappingValues[0]),
			src:         toInt(mappingValues[1]),
			rangeLength: toInt(mappingValues[2]),
		})

		// Last entry
		if i == len(entries)-1 {
			mappers = append(mappers, currentMapper)
		}
	}
	return mappers
}

// for the sake of simplicity/readability
func toInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func getSeeds(s string) []int {
	var seeds []int

	line := strings.Split(s, ":")
	seedsStr := strings.Split(line[1], " ")

	for _, seedStr := range seedsStr {
		if seedStr == "" {
			continue
		}

		seed, err := strconv.Atoi(seedStr)
		if err != nil {
			log.Fatal(err)
		}
		seeds = append(seeds, seed)
	}

	return seeds
}

func getEntries(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var entries []string
	for scanner.Scan() {
		entries = append(entries, scanner.Text())
	}

	return entries
}
