package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	part2()
}

// 87605697
// my main mistake is to use a break (I don't know why) instead of continue
// I use tdd to build the code and it was so satisfying.
func part2() {
	file, err := os.Open("part2-input.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var matrix []string
	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}

	sum := gearRatiosSum(matrix)
	fmt.Printf("Result: %d\n", sum)
}

func gearRatiosSum(matrix []string) int {
	var sum int
	for y := 0; y < len(matrix); y++ {
		// take the first char
		for x := 0; x < len(matrix[y]); x++ {
			// search for *
			if matrix[y][x] != '*' {
				continue
			}

			// Is it a spring ?
			var partNumbers []int
			// checks the same line
			// Before
			if pn := getPartNumberStrBeforeIndex(matrix[y], x); pn != "" {
				partNumbers = append(partNumbers, toInt(pn))
			}
			// After
			if pn := getPartNumberStrAfterIndex(matrix[y], x); pn != "" {
				partNumbers = append(partNumbers, toInt(pn))
			}

			// Above
			if y-1 >= 0 {
				abovePartNumbers := getPartNumbersAroundIdx(matrix[y-1], x)
				if len(partNumbers)+len(abovePartNumbers) > 2 {
					continue
				}
				partNumbers = append(partNumbers, abovePartNumbers...)
			}

			// Below
			if y+1 < len(matrix) {
				abovePartNumbers := getPartNumbersAroundIdx(matrix[y+1], x)
				if len(partNumbers)+len(abovePartNumbers) > 2 {
					continue
				}
				partNumbers = append(partNumbers, abovePartNumbers...)
			}

			// not enough part number, not a gear
			if len(partNumbers) != 2 {
				continue
			}

			sum += partNumbers[0] * partNumbers[1]
		}
	}

	return sum
}

// Wrap atoi function for the sake of visibility
// It's acceptable because we don't really need to manage error
func toInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic("boom")
	}

	return n
}

func getPartNumbersAroundIdx(str string, idx int) []int {
	var partNumbers []int

	before := getPartNumberStrBeforeIndex(str, idx)
	after := getPartNumberStrAfterIndex(str, idx)

	// if the char right above is a digit we have only one part number above.
	if isDigit(str[idx]) {
		return append(partNumbers, toInt(before+string(str[idx])+after))
	}

	if before != "" {
		partNumbers = append(partNumbers, toInt(before))
	}
	if after != "" {
		partNumbers = append(partNumbers, toInt(after))
	}

	return partNumbers
}

func getPartNumberStrBeforeIndex(str string, idx int) string {
	var nStr string
	for offset := 1; idx-offset >= 0; offset++ {
		if !isDigit(str[idx-offset]) {
			break
		}
		nStr = string(str[idx-offset]) + nStr
	}

	return nStr
}

func getPartNumberStrAfterIndex(str string, idx int) string {
	var nStr string
	for offset := 1; idx+offset < len(str); offset++ {
		if !isDigit(str[idx+offset]) {
			break
		}
		nStr += string(str[idx+offset])
	}

	return nStr
}

// result part1 540212
// I've done 2 mistakes
// 1 - does not exclude left check when the number starts the line
// 2 - forgot the numbers that finishes a line
func part1() {
	file, err := os.Open("part1-input.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var matrix []string
	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}

	var sum int
	for y := 0; y < len(matrix); y++ {
		// take the first char
		var nbStr string
		for x := 0; x < len(matrix[y]); x++ {
			// search for number
			if !isDigit(matrix[y][x]) {
				if nbStr == "" {
					continue
				}
				fmt.Println("nbStr:", nbStr)

				// first idx
				first := x - (len(nbStr) + 1)
				if x-1-len(nbStr) < 0 {
					first = x - len(nbStr)
				}
				// last idx
				last := x
				if x+len(nbStr) < len(matrix[y]) {
					last = x + 1
				}

				if isPart(matrix, x, y, first, last) {
					nb, err := strconv.Atoi(nbStr)
					fmt.Printf("%d is a part number\n", nb)
					if err != nil {
						log.Fatal(err)
					}
					sum += nb
				}

				// reset nbStr
				nbStr = ""
				fmt.Println()
				continue
			}
			if isDigit(matrix[y][x]) {
				nbStr += string(matrix[y][x])
			}
		}
		// smell bad, refactor :p
		if nbStr != "" {
			// first idx
			first := len(matrix[y]) - (len(nbStr) + 1)
			if len(matrix[y])-1-len(nbStr) < 0 {
				first = len(matrix[y]) - len(nbStr)
			}
			// last idx
			last := len(matrix[y])
			if len(matrix[y])+len(nbStr) < len(matrix[y]) {
				last = len(matrix[y]) + 1
			}

			if isPart(matrix, len(matrix[y]), y, first, last) {
				nb, err := strconv.Atoi(nbStr)
				fmt.Printf("%d is a part number\n", nb)
				if err != nil {
					log.Fatal(err)
				}
				sum += nb
			}
		}
	}
	fmt.Printf("Result: %d\n", sum)
}

func isPart(matrix []string, x, y, first, last int) bool {
	fmt.Printf("x:%d,y:%d,first:%d,last:%d\n", x, y, first, last)
	// Above.
	if y-1 >= 0 && hasNonDot(matrix[y-1][first:last]) {
		fmt.Printf("Above: something in %q\n", matrix[y-1][first:last])
		return true
	}

	// Below.
	if y+1 < len(matrix) && hasNonDot(matrix[y+1][first:last]) {
		fmt.Printf("Below: something in %q\n", matrix[y+1][first:last])
		return true
	}

	// Current line
	if first != 0 && matrix[y][first] != '.' {
		fmt.Printf("Current line, first letter: something in %q\n", matrix[y][first])
		return true
	}
	if x != last && matrix[y][last-1] != '.' {
		fmt.Printf("Current line, last letter: something in %q\n", matrix[y][last-1])
		return true
	}

	return false
}

func hasNonDot(s string) bool {
	for _, c := range s {
		if c != '.' {
			return true
		}
	}

	return false
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
