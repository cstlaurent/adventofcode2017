package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {

	dat, _ := ioutil.ReadFile("./input.csv")
	s := string(dat[:])

	lines := strings.Split(s, "\n")

	fmt.Printf("Day 2 checksum 1: %d\n", Substractor(lines))
	fmt.Printf("Day 2 checksum 2: %d\n", Divisor(lines))
}

func Substractor(lines []string) int {
	sum := 0
	for _, l := range lines {
		valStrings := strings.Split(l, "\t")
		vals := make([]int, len(valStrings))
		for i := range valStrings {
			vals[i], _ = strconv.Atoi(valStrings[i])
		}
		sort.Ints(vals)
		sum += vals[len(vals)-1] - vals[0]
	}
	return sum
}

func Divisor(lines []string) int {
	sum := 0
	for l := 0; l < len(lines); l++ {
		valStrings := strings.Split(lines[l], "\t")
		vals := make([]int, len(valStrings))
		for i := range valStrings {
			vals[i], _ = strconv.Atoi(valStrings[i])
		}
		sort.Ints(vals)
		found := false
		for i := 0; i < len(vals) && !found; i++ {
			for j := len(vals) - 1; j > i && !found; j-- {
				if vals[j]%vals[i] == 0 {
					sum += vals[j] / vals[i]
					found = true
				}
			}
		}
	}
	return sum
}
