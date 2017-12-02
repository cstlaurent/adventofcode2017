package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

var lines []string

func init() {
	dat, _ := ioutil.ReadFile("./input.csv")
	s := string(dat[:])

	lines = strings.Split(s, "\n")
}

func BenchmarkDivisor(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Divisor(lines)
	}
}
