package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	aoc "github.com/liondadev/aoc2025"
)

// this is the way I originally did part 1 before remembering how math worked
func isRepeatingString(str string, rep string) bool {
	if len(str) == 1 { // satisfy "at least twice" bit
		return false
	}

	s := rep

	if str == s {
		return true
	}

	for len(s) < len(str) {
		s = s + rep
		if s == str {
			return true
		}
	}

	return false
}

func isInvalidProductId(num int) bool {
	numStr := strconv.Itoa(num)
	l := len(numStr)

	for i := 0; i < l/2; i++ {
		part := numStr[:i+1]
		if isRepeatingString(numStr, part) {
			return true
		}
	}

	return false
}

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	ranges := strings.Split(string(input), ",")
	for _, r := range ranges {
		r = strings.TrimSpace(r)

		parts := strings.Split(r, "-")
		lowerStr, upperStr := parts[0], parts[1]
		lower, upper := aoc.MustParseNumber(lowerStr), aoc.MustParseNumber(upperStr)

		for num := lower; num <= upper; num++ {
			if isInvalidProductId(num) {
				sum += num
			}
		}
	}

	fmt.Println(sum)
}
