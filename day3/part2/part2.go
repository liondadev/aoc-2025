package main

import (
	"fmt"

	aoc "github.com/liondadev/aoc2025"
)

// getMaxJoltage gets the biggest k-digit number that can be made from
// the digits in order.
func getMaxJoltage(rack string, k int) int {

}

func main() {
	fmt.Println(getMaxJoltage("811111111111119")) // 811111111119
	//input, err := os.ReadFile("./input.txt")
	//if err != nil {
	//	panic(err)
	//}
	//
	//sum := 0
	//sc := bufio.NewScanner(bytes.NewBuffer(input))
	//sc.Split(bufio.ScanLines)
	//for sc.Scan() {
	//	line := sc.Text()
	//	line = strings.TrimSpace(line)
	//	if line == "" {
	//		continue
	//	}
	//
	//	k := getMaxJoltage(line)
	//	sum += k
	//}
	//
	//fmt.Println(sum)
}
