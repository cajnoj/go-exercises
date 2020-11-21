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
	//fmt.Println(s)

	if s[0] == '-' || s[0] == '+' {
		b.WriteByte(s[0])
		s = s[1:]
	}
	//fmt.Println(s)

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		commaInt(b, s[:dot])
		b.WriteString(s[dot:])
		//fmt.Println(dot, s[:dot], s[dot:], b.String())
	} else {
		commaInt(b, s)
		//fmt.Println(dot, s, b.String())
	}

	return b.String()
}

// comma inserts commas in a non-negative decimal integer string.
func commaInt(b *bytes.Buffer, s string) {
	n := len(s)
	//println("ci", s, n, b.String())
	for i := 0; i < n; i++ {
		if i > 0 && (n-i)%3 == 0 {
			b.WriteByte(',')
			//fmt.Println("ci", ',', b.String())
		}
		b.WriteByte(s[i])
		//fmt.Println("ci", s[i], b.String())
	}
}
