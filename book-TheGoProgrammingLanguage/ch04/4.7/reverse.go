package main

import (
	"fmt"
	"unicode/utf8"
)

// reverse reverses a slice of ints in place.
func reverse(s []byte) []byte {
	r, size := utf8.DecodeLastRune(s)

	if size == 0 {
		return s
	}

	lastChar := len(s) - size

	for i := lastChar; i > 0; i -= size {
		r, size = utf8.DecodeLastRune(s[:i])
		// The last iteration will not append anything
		s = append(s, []byte(string(r))...)
	}

	return s[lastChar:]
}

func main() {
	a := []byte("אבגד הוזח")
	fmt.Printf("%v %q len %d cap %d\n", a, string(a), len(a), cap(a))
	a = reverse(a)
	fmt.Printf("%v %q len %d cap %d\n", a, string(a), len(a), cap(a))
}
