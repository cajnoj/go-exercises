package main

import (
	"crypto/sha256"
	"fmt"
)

const (
	SHA256Bytes = 32
)

type sha256Hash [SHA256Bytes]uint8

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	diffBits := bitCountDiff256(c1, c2)
	fmt.Printf("c1 %x\nc2 %x\ndiffBits %d\n", c1, c2, diffBits)
}

func bitCountDiff256(b1, b2 sha256Hash) int {
	// Count bits in difference
	var result int
	for i := 0; i < SHA256Bytes; i++ {
		result += int(pc[b1[i]^b2[i]])
	}
	return result
}

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
