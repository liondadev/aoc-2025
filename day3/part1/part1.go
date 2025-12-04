package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	aoc "github.com/liondadev/aoc2025"
)

// getMaxJoltage gets the maximum joltage for an integer.
func getMaxJoltage(num int) int {
	first := num / 10 % 10
	second := num % 10

	n := num
	for n > 0 {
		nSecond := n / 10 % 10
		n /= 10

		nFirst := n / 10 % 10
		n /= 10

		if nSecond > second {
			second = nSecond
		}

		if nFirst > first {
			first = nFirst
		}
	}

	return first*10 + second
}

func getMaxJoltageString(rack string) int {
	m := 0

	for i := 0; i < len(rack)-1; i++ {
		for j := i + 1; j < len(rack); j++ {
			str := string(rack[i]) + string(rack[j])
			num := aoc.MustParseNumber(str)
			if num > m {
				m = num
			}

		}
	}

	return m
}

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	sc := bufio.NewScanner(bytes.NewBuffer(input))
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		line := sc.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		k := getMaxJoltageString(line)
		sum += k
	}

	fmt.Println(sum)
}
