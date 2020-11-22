package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Error: Expecting 2 args\n")
		os.Exit(1)
	}

	a, b := os.Args[1], os.Args[2]

	not := "not "
	if anagrams(a, b) {
		not = ""
	}

	fmt.Printf("%q and %q are %sanagrams\n", a, b, not)
}

func anagrams(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	ra, rb := []rune(a), []rune(b)

	if len(ra) != len(rb) {
		return false
	}

	for i := 0; i < len(ra); i++ {
		if ra[i] != rb[len(ra)-i-1] {
			return false
		}
	}

	return true
}
