package main

import (
	"testing"
)

var lines [][]int

func init() {
	lines = readCsv()
}

func BenchmarkSubstractor(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		substractor(lines)
	}
}
func BenchmarkDivisor(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		divisor(lines)
	}
}
