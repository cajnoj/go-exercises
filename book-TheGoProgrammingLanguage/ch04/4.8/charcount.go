// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode"
)

func main() {
	counters := make(map[string]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		switch {
		case r == unicode.ReplacementChar && n == 1:
			counters["invalid"]++
		case unicode.IsLetter(r):
			counters["letter"]++
		case unicode.IsDigit(r):
			counters["digit"]++
		}
	}

	for _, k := range getKeysSorted(counters) {
		fmt.Printf("%15d\t%ss\n", counters[k], k)
	}
}

func getKeysSorted(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}
