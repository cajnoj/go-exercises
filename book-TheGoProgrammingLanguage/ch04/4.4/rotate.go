package main

import (
	"fmt"
	"os"
	"time"
)

const arrayLength = 24

var initialArray []int

// Rotate s by r
func rotate1(s []int, r int) []int {
	if r > 0 {
		r = r%len(s) - len(s)
	} else if r == 0 {
		return s
	} else {
		r = r % len(s)
	}

	s = append(s, s[:-r]...)
	s = s[-r:]

	return s
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func next(i, jump, size int) int {
	if i < jump {
		i += size
	}

	return i - jump
}

func rotateCycle(s []int, start int, jump int) {
	for i := start; next(i, jump, len(s)) != start; i = next(i, jump, len(s)) {
		s[i], s[next(i, jump, len(s))] = s[next(i, jump, len(s))], s[i]
	}
}

// Rotate s by r
func rotate2(s []int, r int) []int {
	if r < 0 {
		r = r%len(s) + len(s)
	} else if r == 0 {
		return s
	} else {
		r = r % len(s)
	}

	g := gcd(r, len(s))
	if g == 1 {
		// Single cycle
		rotateCycle(s, 0, r)
	} else {
		// g cycles
		for m := 0; m < g; m++ {
			rotateCycle(s, m, r)
		}
	}

	return s
}

func exitWhenDifferent(s1, s2 []int) {
	for j := 0; j < arrayLength; j++ {
		if s1[j] != s2[j] {
			fmt.Fprintln(os.Stderr, "Difference!")
			os.Exit(1)
		}
	}
}

func init() {
	initialArray = make([]int, arrayLength)
	for i := 0; i < arrayLength; i++ {
		initialArray[i] = i
	}
}

func main() {
	var s1, s2 []int
	s1 = make([]int, arrayLength)
	s2 = make([]int, arrayLength)
	for i := -arrayLength; i <= arrayLength; i++ {
		copy(s1, initialArray)
		copy(s2, initialArray)
		s1 = rotate1(s1, i)
		s2 = rotate1(s2, i)
		fmt.Printf("%v >> %2d == %v\n", initialArray, i, s1)
		exitWhenDifferent(s1, s2)
	}
	//
	fmt.Println("Measuring rotate1")
	start := time.Now()
	for n := 0; n < 1e5; n++ {
		for i := -arrayLength; i <= arrayLength; i++ {
			copy(s1, initialArray)
			s1 = rotate1(s1, i)
		}
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	//
	fmt.Println("Measuring rotate2")
	start = time.Now()
	for n := 0; n < 1e5; n++ {
		for i := -arrayLength; i <= arrayLength; i++ {
			copy(s2, initialArray)
			s1 = rotate2(s2, i)
		}
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
