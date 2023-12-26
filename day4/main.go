package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// part1: 21558
// part2: 10425665
func main() {
	entries := getEntries("part2-input.txt")
	results := getCardsNumber(entries)
	fmt.Println(getTotalCount(results))
}

func getTotalCount(results []int) int {
	var total int
	for _, r := range results {
		total += r
	}

	return total
}

func getCardsNumber(entries []string) []int {
	max := len(entries)
	results := make([]int, max)

	for i := 0; i < max; i++ {
		nb := winingCardsNb(entries[i])
		results[i]++

		for j := i + 1; j <= i+nb && j < max; j++ {
			results[j] += results[i]
		}
	}

	return results
}

func score(winingCardsNb int) int {
	if winingCardsNb == 0 {
		return 0
	}

	return 1 << (winingCardsNb - 1)
}

func winingCardsNb(s string) int {
	line := strings.Split(s, ":")
	values := strings.Split(line[1], "|")
	winning := strings.Split(values[0], " ")

	hasWon := make(map[string]struct{}, len(winning))
	for _, v := range winning {
		if v != "" {
			hasWon[v] = struct{}{}
		}
	}

	var nb int
	actual := strings.Split(values[1], " ")
	for _, v := range actual {
		if _, ok := hasWon[v]; !ok {
			continue
		}
		nb++
	}

	return nb
}

func getEntries(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var matrix []string
	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}

	return matrix
}
