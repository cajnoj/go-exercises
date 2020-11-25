package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	var s []byte = []byte("אבג\r\t\n דהו  זחט      ")
	fmt.Printf("%v\t%q\n", s, string(s))
	s = squash(s)
	fmt.Printf("%v\t%q\n", s, string(s))
}

func squash(s []byte) []byte {
	const noSpace = -1
	var spaceStart int = noSpace
	var spaceSize int

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRune(s[i:])

		if unicode.IsSpace(r) {
			if spaceStart == noSpace {
				spaceStart = i
				spaceSize = size
			} else {
				spaceSize += size
			}
			i += size
		} else {
			if spaceStart == noSpace {
				i += size
			} else {
				// Space ended
				s2 := s[spaceStart+spaceSize:]
				s = append(s[:spaceStart], 0x20)
				s = append(s, s2...)
				spaceStart = noSpace
			}
		}
	}

	if spaceStart != noSpace {
		// String ended with space
		s = s[:spaceStart]
	}

	return s
}
