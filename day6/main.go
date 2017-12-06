package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("./input.txt")
	s := string(dat[:])

	lines := strings.Split(s, "\r\n")

	fmt.Printf("Day 6 part 1: %d steps\n", calc(lines))
	fmt.Printf("Day 6 part 2: %d steps\n", calc(lines))
}

func calc(lines []string) int {
	return 4
}
