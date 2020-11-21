package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	n := 1
	for i := 0; i < 6; i++ {
		n *= 10
		s := strconv.Itoa(n)
		fmt.Printf("%10s %10s\n", s, comma(s))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	var b bytes.Buffer
	for i := 0; i < n; i++ {
		if i > 0 && (n-i)%3 == 0 {
			b.WriteByte(',')
		}
		b.WriteByte(s[i])
	}
	return b.String()
}
