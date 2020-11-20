// Package c64 implements Mandelbrot with complex64
package c64

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
		y := float32(py)/float32(height)*float32(ymax-ymin) + float32(ymin)
		for px := 0; px < width; px++ {
			x := float32(px)/float32(width)*float32(xmax-xmin) + float32(xmin)
			z := complex64(complex(x, y))
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func abs(z complex64) float32 {
	return float32(math.Sqrt(float64(real(z)*real(z) + imag(z)*imag(z))))
}

func mandelbrot(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
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
