package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Measuring string concatenation method")
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Printf("length %d %.2fs elapsed\n", len(s), time.Since(start).Seconds())
	//
	fmt.Println("Measuring strings.Join() method")
	start = time.Now()
	s = strings.Join(os.Args[1:], " ")
	fmt.Printf("length %d %.2fs elapsed\n", len(s), time.Since(start).Seconds())
}
