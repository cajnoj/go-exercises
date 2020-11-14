// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileCounts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "Stdin", counts, fileCounts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts, fileCounts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)

			for fn, m := range fileCounts[line] {
				fmt.Printf("...\t%d\t%s\n", m, fn)
			}
		}
	}
}

func countLines(f *os.File, fn string, counts map[string]int, fileCounts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if counts[line] == 0 {
			fileCounts[line] = make(map[string]int)
		}
		counts[line]++
		fileCounts[line][fn]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
