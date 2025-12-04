package main

import (
	"fmt"
	"os"
	"strings"
)

const PaperRoll = '@'

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(input), "\n")

	accessible := 0
	for row := 0; row < len(rows); row++ {
		if rows[row] == "" {
			continue
		}

		rows[row] = strings.TrimSpace(rows[row]) // trim newline and stuff
		for col := 0; col < len(rows[row]); col++ {
			adj := 0 // # of adjacent rows

			if rows[row][col] != PaperRoll {
				fmt.Print(".")
				continue
			}

			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if dx == 0 && dy == 0 {
						continue // don't count itself
					}

					if row+dx < 0 || row+dx > len(rows)-1 {
						continue
					}
					if col+dy < 0 || col+dy > len(rows[row])-1 {
						continue
					}

					if rows[row+dx][col+dy] == PaperRoll {
						adj++
					}
				}
			}

			fmt.Print(adj)
			if adj < 4 {
				accessible++
			}
		}

		fmt.Println("")
	}

	fmt.Println(accessible)
}
