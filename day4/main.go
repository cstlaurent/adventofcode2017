package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("./input.txt")
	s := string(dat[:])

	lines := strings.Split(s, "\n")
	fmt.Printf("Day 4 part 1: %d valid passphrases\n", validPass(lines))
	fmt.Printf("Day 4 part 2: %d valid passphrases\n", validPassAnagram(lines))
}

func validPass(lines []string) int {
	valid := 0
	for _, line := range lines {
		words := strings.Split(line, " ")
		found := false
		for i := 0; i < len(words) && !found; i++ {
			for j := len(words) - 1; j > i && !found; j-- {
				if words[i] == words[j] {
					found = true
				}
			}
		}
		if !found {
			valid++
		}
	}
	return valid
}

func validPassAnagram(lines []string) int {
	valid := 0
	for _, line := range lines {
		words := strings.Split(line, " ")
		found := false
		for i := 0; i < len(words) && !found; i++ {
			for j := len(words) - 1; j > i && !found; j-- {
				if sortWord(words[i]) == sortWord(words[j]) {
					found = true
				}
			}
		}
		if !found {
			valid++
		}
	}
	return valid
}

func sortWord(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
