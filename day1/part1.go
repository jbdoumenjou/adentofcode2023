package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func part1() {
	file, err := os.Open("part1-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result int

	for scanner.Scan() {
		content := scanner.Text()
		result += getSum(content)
	}

	fmt.Println("Result: ", result)
}

func getSum(s string) int {
	if len(s) == 0 || s == "\n" {
		return 0
	}

	firstIdx := strings.IndexFunc(s, unicode.IsDigit)
	first := int(s[firstIdx] - '0')

	lastIdx := strings.LastIndexFunc(s, unicode.IsDigit)
	last := int(s[lastIdx] - '0')

	return first*10 + last
}
