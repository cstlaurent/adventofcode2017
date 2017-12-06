package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := [...]int{0, 5, 10, 0, 11, 14, 13, 4, 11, 8, 8, 7, 1, 4, 12, 11}
	s1, s2 := calc(input)
	fmt.Printf("Day 6 part 1: seen at %d and seen again at %d \n", s1, s2)
}

func calc(input [16]int) (int, int) {
	set := newStringSet()
	seen := false
	seen2 := false
	cycle := 0
	cycle2 := 0
	for !seen2 {
		lv := 0
		li := 0
		// find largest
		for i, v := range input {
			if v > lv {
				li = i
				lv = v
			}
		}

		// Reassign bank
		input[li] = 0
		for index := 1; index <= lv; index++ {
			input[(li+index)%len(input)]++
		}
		stringVal := intToString(input)
		if set.get(stringVal) {
			if seen {
				seen2 = true
			} else {
				seen = true
			}
		} else {
			set.add(stringVal)
		}
		cycle++
		if seen {
			cycle2++
		}
	}
	return cycle, cycle2
}

func intToString(a [16]int) string {
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}

	return strings.Join(b, ",")
}

type stringSet struct {
	set map[string]bool
}

func newStringSet() *stringSet {
	return &stringSet{make(map[string]bool)}
}

func (set *stringSet) add(i string) bool {
	_, found := set.set[i]
	set.set[i] = true
	return !found //False if it existed already
}

func (set *stringSet) get(i string) bool {
	_, found := set.set[i]
	return found //true if it existed already
}
