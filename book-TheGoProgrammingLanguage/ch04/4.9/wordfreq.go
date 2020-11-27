package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	freqs := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		freqs[input.Text()]++
	}

	for k, v := range freqs {
		fmt.Println(k, v)
	}
}
