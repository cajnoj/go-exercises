// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 10240, 10240
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, superMandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) (uint8, uint8, uint8) {
	const iterations = 200

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if (real(v) > 2 || imag(v) > 2) && cmplx.Abs(v) > 2 {
			return uint8(math.Pow(1.028, float64(n))), 0, 0
		}
	}
	return 0, 0, 0
}

func superMandelbrot(z complex128) color.Color {
	const (
		xHalf = float64(1) / width / 2.0 * (xmax - xmin)
		yHalf = float64(1) / height / 2.0 * (ymax - ymin)
	)
	//
	r1, g1, b1 := mandelbrot(z)
	r2, g2, b2 := mandelbrot(z + complex(xHalf, 0))
	r3, g3, b3 := mandelbrot(z + complex(0, yHalf))
	r4, g4, b4 := mandelbrot(z + complex(xHalf, yHalf))
	//
	r := (uint16(r1) + uint16(r2) + uint16(r3) + uint16(r4)) / 4
	g := (uint16(g1) + uint16(g2) + uint16(g3) + uint16(g4)) / 4
	b := (uint16(b1) + uint16(b2) + uint16(b3) + uint16(b4)) / 4
	//
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}
