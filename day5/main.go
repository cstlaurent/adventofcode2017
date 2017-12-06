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
	fmt.Printf("Day 5 part 1: %d steps\n", calcJumps(lines))
}

func calcJumps(lines []string) int {
	fmt.Println(len(lines))
	jumps := make([]int, len(lines))
	for i, val := range lines {
		integ, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		jumps[i] = integ
	}
	currentVal := 0
	oldVal := 0
	jumpsTaken := 0
	for {
		// fmt.Println(currentVal)
		oldVal = currentVal
		currentVal = currentVal + jumps[currentVal]

		jumpsTaken++
		if currentVal < 0 || currentVal >= len(lines) {
			break
		}
		jumps[oldVal]++
	}
	return jumpsTaken
}
