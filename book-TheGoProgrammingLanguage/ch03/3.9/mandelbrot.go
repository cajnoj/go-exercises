// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"book-TheGoProgrammingLanguage/ch03/3.9/bfloat"
	"book-TheGoProgrammingLanguage/ch03/3.9/brat"
	"book-TheGoProgrammingLanguage/ch03/3.9/c128"
	"book-TheGoProgrammingLanguage/ch03/3.9/c64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func create(fn string) *os.File {
	var f *os.File
	var err error

	if f, err = os.Create(fn); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return f
}

func main() {
	http.HandleFunc("/", handler)
	log.Print("Server running.")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type drawFunc func(io.Writer, float64, float64, float64, float64, int, int)

func parseFloat(s string, def float64) float64 {
	var f float64
	var err error
	if f, err = strconv.ParseFloat(s, 64); err != nil {
		log.Print(err)
		return def
	}
	return f
}

func parseInt(s string, def int) int {
	var i int
	var err error
	if i, err = strconv.Atoi(s); err != nil {
		log.Print(err)
		return def
	}
	return i
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Defaults
	var defaultMethod drawFunc = c128.Draw
	const (
		// defaultXloc   = 0.40052917
		// defaultYloc   = 0.3505417
		// defaultDelta  = 4.17e-06
		// defaultPixels = 512
		defaultXloc   = -2.0
		defaultYloc   = -2.0
		defaultDelta  = 4.0
		defaultPixels = 1024
	)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
		return
	}

	var method drawFunc
	var xloc, yloc, delta float64
	var pixels int

	method, xloc, yloc, delta, pixels = defaultMethod, defaultXloc, defaultYloc, defaultDelta, defaultPixels

	for k, v := range r.Form {
		switch k {
		case "method":
			switch v[0] {
			case "complex64":
				method = c64.Draw
			case "complex128":
				method = c128.Draw
			case "big.Float":
				method = bfloat.Draw
			case "big.Rat":
				method = brat.Draw
			default:
				log.Printf("Invalid method %s\n", v[0])
			}
		case "xloc":
			xloc = parseFloat(v[0], defaultXloc)
		case "yloc":
			yloc = parseFloat(v[0], defaultYloc)
		case "delta":
			delta = parseFloat(v[0], defaultDelta)
		case "pixels":
			pixels = parseInt(v[0], defaultPixels)
		default:
			log.Printf("Invalid key %s\n", k)
		}

	}

	log.Printf("Request=%s;xloc=%g;yloc=%g;delta=%g;pixels=%d\n", r.Form.Encode(), xloc, yloc, delta, pixels)
	method(w, xloc, yloc, xloc+delta, yloc+delta, pixels, pixels)
}
