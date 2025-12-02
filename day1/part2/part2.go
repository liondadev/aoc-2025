package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	timesZero := 0
	current := 50
	sc := bufio.NewScanner(bytes.NewBuffer(input))
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		line := sc.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		dir := line[0:1]
		clicksStr := line[1:]
		clicksNum, err := strconv.Atoi(clicksStr)
		if err != nil {
			fmt.Printf("error parsing clicks number '%s': %s\n", clicksStr, err.Error())
			continue
		}

		sign := +1
		if dir == "L" {
			sign = -1
		}

		// moves by one
		step := func() {
			newNotInRange := current + sign
			newCurrent := (newNotInRange%100 + 100) % 100 // thanks go for not having true modulo :)
			//fmt.Printf("%d : %s -> %d\n", current, line, newCurrent)
			current = newCurrent

			if current == 0 {
				timesZero++
			}
		}

		for i := 0; i < clicksNum; i++ {
			step()
		}

	}

	fmt.Printf("Times hit / passed zero: %d\n", timesZero)
}
