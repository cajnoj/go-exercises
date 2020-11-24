package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

type bitCountFunc func(string, string) int

func main() {
	s1, s2 := "x", "X"

	var fun bitCountFunc
	var algo string

	switch len(os.Args) {
	case 1:
		fun = bitCountDiff256
	case 2:
		switch os.Args[1] {
		case "SHA384":
			algo = os.Args[1]
			fun = bitCountDiff384
		case "SHA512":
			algo = os.Args[1]
			fun = bitCountDiff512
		default:
			algo = "SHA256"
			fun = bitCountDiff256
		}
	default:
		fmt.Fprintf(os.Stderr, "Error: No more than 1 argument allowed\n")
		os.Exit(1)
	}

	diffBits := fun(s1, s2)

	fmt.Printf("s1 %q\ts2 %q\tdiffBits %d\thash %s\n", s1, s2, diffBits, algo)
}

type hashFunc func([]byte)

func bitCountDiff256(s1, s2 string) int {
	b1 := sha256.Sum256([]byte(s1))
	b2 := sha256.Sum256([]byte(s2))

	var result int
	for i := 0; i < sha256.Size; i++ {
		result += diffBits(b1[i], b2[i])
	}
	return result
}

func bitCountDiff384(s1, s2 string) int {
	b1 := sha512.Sum384([]byte(s1))
	b2 := sha512.Sum384([]byte(s2))

	var result int
	for i := 0; i < sha512.Size384; i++ {
		result += diffBits(b1[i], b2[i])
	}
	return result
}

func bitCountDiff512(s1, s2 string) int {
	b1 := sha512.Sum512([]byte(s1))
	b2 := sha512.Sum512([]byte(s2))

	var result int
	for i := 0; i < sha512.Size; i++ {
		result += diffBits(b1[i], b2[i])
	}
	return result
}

// pc[i] is the population count of i.
var pc [256]byte

func diffBits(u1, u2 uint8) int {
	return int(pc[u1^u2])
}

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
