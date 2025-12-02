package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	aoc "github.com/liondadev/aoc2025"
)

func digitsOf(num int) int {
	return int(math.Floor(math.Log10(float64(num))) + 1)
}

func isInvalidProductId(num int) bool {
	digits := digitsOf(num)
	halfDigits := digits / 2

	part := num / int(math.Pow(10, float64(halfDigits)))
	uPart := part * int(math.Pow(10, float64(halfDigits)))
	return uPart+part == num
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
