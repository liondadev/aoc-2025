package main

import (
	"fmt"
	"os"
	"strings"
)

const PaperRoll = '@'

var shouldPrint = len(os.Args) < 2

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(input), "\n")

	removePossible := func() int {
		accessible := 0
		for row := 0; row < len(rows); row++ {
			if rows[row] == "" {
				continue
			}

			rows[row] = strings.TrimSpace(rows[row]) // trim newline and stuff
			for col := 0; col < len(rows[row]); col++ {
				adj := 0 // # of adjacent rows

				if rows[row][col] != PaperRoll {
					if shouldPrint {
						fmt.Print(".")
					}
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

				if shouldPrint {
					fmt.Print(adj)
				}
				if adj < 4 {
					accessible++
					rows[row] = rows[row][:col] + "x" + rows[row][col+1:]
				}
			}

			if shouldPrint {
				fmt.Println("")
			}
		}

		return accessible
	}

	removedSoFar := 0
	tickNum := 0
	for {
		if shouldPrint {
			fmt.Printf("\ntick %d:\n", tickNum)
		}
		removed := removePossible()
		removedSoFar += removed
		tickNum++

		// we're done :)
		if removed == 0 {
			break
		}
	}

	fmt.Println(removedSoFar)
}
