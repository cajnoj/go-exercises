package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	f := -0.001
	for i := 0; i < 9; i++ {
		s := strconv.FormatFloat(f, 'f', -1, 64)
		fmt.Printf("%20s %20s\n", s, comma(s))
		f *= -1.0 * 10.0
	}
}

func comma(s string) string {
	var b *bytes.Buffer = new(bytes.Buffer)

	if s[0] == '-' || s[0] == '+' {
		b.WriteByte(s[0])
		s = s[1:]
	}

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		commaInt(b, s[:dot])
		b.WriteString(s[dot:])
	} else {
		commaInt(b, s)
	}

	return b.String()
}

// comma inserts commas in a non-negative decimal integer string.
func commaInt(b *bytes.Buffer, s string) {
	n := len(s)

	for i := 0; i < n; i++ {
		if i > 0 && (n-i)%3 == 0 {
			b.WriteByte(',')
		}
		b.WriteByte(s[i])
	}
}
