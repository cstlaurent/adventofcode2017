package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("./input.txt")
	s := string(dat[:])

	lines := strings.Split(s, "\r\n")
	fmt.Printf("Day 5 part 1: %d steps\n", calcJumps(lines, false))
	fmt.Printf("Day 5 part 2: %d steps\n", calcJumps(lines, true))
}

func calcJumps(lines []string, day2 bool) int {
	jumps := make([]int, len(lines))
	for i, val := range lines {
		integ, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		jumps[i] = integ
	}
	current := 0
	steps := 0
	for current < len(jumps) && current >= 0 {
		steps++
		offset := jumps[current]
		if day2 && offset >= 3 {
			jumps[current]--
		} else {
			jumps[current]++
		}
		current += offset
	}
	return steps
}
