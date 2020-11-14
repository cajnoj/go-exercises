package main

import (
	"book-TheGoProgrammingLanguage/ch02/2.3/popcount"
	"fmt"
	"time"
)

const (
	iterations = 100000000
)

func main() {
	measure("PopCount1", popcount.PopCount1)
	measure("PopCount2", popcount.PopCount2)
}

func measure(s string, f func(uint64) int) {
	var x uint64

	start := time.Now()
	for j := 0; j < iterations; j++ {
		f(x)
	}
	fmt.Printf("%s %.2fs elapsed\n", s, time.Since(start).Seconds())
}
