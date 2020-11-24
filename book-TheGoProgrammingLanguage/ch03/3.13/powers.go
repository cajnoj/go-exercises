package main

import (
	"fmt"
)

const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	// Following constants would cause an overflow
	// ZB = 1000 * EB
	// YB = 1000 * ZB
)

func main() {
	fmt.Printf("KB%25d\n", KB)
	fmt.Printf("MB%25d\n", MB)
	fmt.Printf("GB%25d\n", GB)
	fmt.Printf("TB%25d\n", TB)
	fmt.Printf("PB%25d\n", PB)
	fmt.Printf("EB%25d\n", EB)
}
