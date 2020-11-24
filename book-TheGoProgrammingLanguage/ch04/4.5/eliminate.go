package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args[1:]
	fmt.Printf("Before: %v\n", s)
	s = elimiate(s)
	fmt.Printf("After:  %v\n", s)
}

func elimiate(s []string) []string {
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}

	return s
}
