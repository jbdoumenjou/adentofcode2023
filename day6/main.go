package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// part1: 5133600
// part2: 40651271
func main() {
	entries := getEntries("part1-input.txt")
	r := getRace(entries)

	fmt.Println(r.nbOfWaysToBeatDistance())
}

func getRace(entries []string) race {
	return race{
		Time:     getValue(entries[0]),
		Distance: getValue(entries[1]),
	}
}

func getValue(s string) int {
	line := strings.Split(s, ":")
	valueStr := strings.ReplaceAll(line[1], " ", "")

	return toInt(valueStr)
}

func part1() {
	entries := getEntries("part1-input.txt")
	races := getRaces(entries)

	fmt.Println(getResult(races))
}

func getResult(races []race) int {
	var result = 1

	for _, race := range races {
		result *= race.nbOfWaysToBeatDistance()
	}

	return result
}

type race struct {
	Time     int
	Distance int
}

func (r race) nbOfWaysToBeatDistance() int {
	var nbWays int

	for i := 1; i < r.Time-1; i++ {
		if (r.Time-i)*i > r.Distance {
			nbWays++
		}
	}

	return nbWays
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

func getRaces(entries []string) []race {
	times := getValues(entries[0])
	distances := getValues(entries[1])

	var races []race
	for i := 0; i < len(times); i++ {
		races = append(races, race{
			Time:     times[i],
			Distance: distances[i],
		})
	}

	return races
}

func getValues(s string) []int {
	var times []int

	line := strings.Split(s, ":")
	values := strings.Split(line[1], " ")
	for _, value := range values {
		if value == "" {
			continue
		}
		times = append(times, toInt(value))
	}

	return times
}

func toInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return val
}
