package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("part1-input.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var result int
	for scanner.Scan() {
		result += score(scanner.Text())
	}
	fmt.Println(result)
}

func score(s string) int {
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

	if nb == 0 {
		return 0
	}

	return 1 << (nb - 1)
}
