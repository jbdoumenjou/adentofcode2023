package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1() {
	file, err := os.Open("part1-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	bag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	scanner := bufio.NewScanner(file)
	i := 1
	var sum int
	for scanner.Scan() {
		content := scanner.Text()
		if isPossibleGame(content, bag) {
			sum += i
		}

		i++
	}
	fmt.Println(sum)
}

func isPossibleGame(s string, bag map[string]int) bool {
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	game := strings.Split(s, ":")
	// 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	sets := strings.Split(game[1], ";")
	for _, set := range sets {
		// 3 blue, 4 red
		cubes := strings.Split(set, ",")
		for _, cube := range cubes {
			// 3 blue
			elt := strings.Split(strings.TrimSpace(cube), " ")
			val, err := strconv.Atoi(elt[0])
			if err != nil {
				log.Fatal(err)
			}

			if val > bag[elt[1]] {
				return false
			}
		}
	}

	return true
}
