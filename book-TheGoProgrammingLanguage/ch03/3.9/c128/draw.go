// Package c128 implements Mandelbrot with complex64
package c128

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
)

// Draw function
func Draw(w io.Writer, xmin, ymin, xmax, ymax float64, width, height int) {

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex128(complex(x, y))
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func abs(z complex128) float64 {
	return math.Sqrt(real(z)*real(z) + imag(z)*imag(z))
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if abs(v) > 2 {
			val := uint8(255 - contrast*n)
			R := uint8(val & 0xab)
			G := uint8(val & 0xcd)
			B := uint8(val & 0xef)
			A := uint8(255)
			return color.RGBA{R, G, B, A}
		}
	}
	return color.Black
}
