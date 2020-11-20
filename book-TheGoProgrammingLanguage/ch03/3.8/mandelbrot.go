// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"book-TheGoProgrammingLanguage/ch03/3.8/bfloat"
	"book-TheGoProgrammingLanguage/ch03/3.8/brat"
	"book-TheGoProgrammingLanguage/ch03/3.8/c128"
	"book-TheGoProgrammingLanguage/ch03/3.8/c64"
	"fmt"
	"io"
	"os"
	"strconv"
)

type drawFunc func(io.Writer, float64, float64, float64, float64, int, int)

func create(fn string) *os.File {
	var f *os.File
	var err error

	if f, err = os.Create(fn); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return f
}

func printUsage() {
	fmt.Printf("Usage: %s complex64|complex128|big.Float|big.Rat xloc yloc delta pixels\n", os.Args[0])
	fmt.Println("Examples:")
	fmt.Printf("\t%s complex128 -2 -2 4 256\n", os.Args[0])
	fmt.Printf("\t%s big.Float 0.40052917 0.3505417 4.17e-06 128\n", os.Args[0])
}

func getFloatArg(n int) float64 {
	var f float64
	var err error
	if f, err = strconv.ParseFloat(os.Args[n], 64); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(n)
	}
	return f
}

func getIntArg(n int) int {
	var i int
	var err error
	if i, err = strconv.Atoi(os.Args[n]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(n)
	}
	return i
}

func getDrawFunc(methodName string) drawFunc {
	var method drawFunc
	switch methodName {
	case "complex64":
		method = c64.Draw
	case "complex128":
		method = c128.Draw
	case "big.Float":
		method = bfloat.Draw
	case "big.Rat":
		method = brat.Draw
	default:
		printUsage()
		os.Exit(11)
	}
	return method
}

func getParams() (drawFunc, float64, float64, float64, int) {
	if len(os.Args) != 6 {
		printUsage()
		os.Exit(10)
	}

	return getDrawFunc(os.Args[1]), getFloatArg(2), getFloatArg(3), getFloatArg(4), getIntArg(5)
}

func main() {
	method, xloc, yloc, delta, pixels := getParams()

	method(create("out.png"), xloc, yloc, xloc+delta, yloc+delta, pixels, pixels)
}
