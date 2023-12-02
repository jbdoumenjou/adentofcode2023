package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	part2()
}

func part2() {
	//file, err := os.Open("part2-example.txt")
	file, err := os.Open("part2-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	keys := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9}
	var result int

	for scanner.Scan() {
		content := scanner.Text()
		sum := getSum2(content, keys)
		fmt.Println(sum)
		result += sum
	}

	fmt.Println("Result: ", result)
}

func getSum2(s string, keys map[string]int) int {
	if len(s) == 0 || s == "\n" {
		return 0
	}

	firstIdx, firstVal, lastIdx, lastVal := -1, -1, -1, -1
	for key, i := range keys {
		if index := strings.Index(s, key); index != -1 {
			if firstIdx == -1 || index < firstIdx {
				firstIdx = index
				firstVal = i
			}
		}
		if index := strings.LastIndex(s, key); index != -1 {
			if lastIdx == -1 || index > lastIdx {
				lastIdx = index
				lastVal = i
			}
		}

	}

	return firstVal*10 + lastVal
}
