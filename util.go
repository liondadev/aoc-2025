package aoc

import "strconv"

func MustParseNumber(n string) int {
	in, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}

	return in
}
