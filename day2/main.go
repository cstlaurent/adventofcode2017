package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	vals := readCsv()
	// fmt.Printf("Day 2 checksum 1: %d\n", substractor(vals))
	fmt.Printf("Day 2 checksum 2: %d\n", divisor(vals))
}

func readCsv() [][]int {
	dat, _ := ioutil.ReadFile("./input.csv")
	s := string(dat[:])

	lines := strings.Split(s, "\n")

	vals := make([][]int, len(lines))

	for l := 0; l < len(lines); l++ {
		valStrings := strings.Split(lines[l], "\t")
		line := make([]int, len(valStrings))
		for i := range valStrings {
			line[i], _ = strconv.Atoi(valStrings[i])
		}
		vals[l] = line
	}
	return vals
}

func substractor(lines [][]int) int {
	sum := 0
	for l := 0; l < len(lines); l++ {
		min := 999999999
		max := 0
		lineData := lines[l]
		for c := 0; c < len(lineData); c++ {
			if min > lineData[c] {
				min = lineData[c]
			}
			if max < lineData[c] {
				max = lineData[c]
			}
		}
		sum += max - min
	}
	return sum
}

func divisor(lines [][]int) int {
	sum := 0
	for l := 0; l < len(lines); l++ {
		lineData := lines[l]
		sort.Ints(lineData)
		found := false
		for i := 0; i < len(lineData) && !found; i++ {
			for j := len(lineData) - 1; j > i && !found; j-- {
				if lineData[j]%lineData[i] == 0 {
					sum += lineData[j] / lineData[i]
					found = true
				}
			}
		}
	}
	return sum
}
