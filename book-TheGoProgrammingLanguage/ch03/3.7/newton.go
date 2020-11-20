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

func main() {
	const (
		xmin, ymin, xmax, ymax = -1.4, -1.4, 1.4, 1.4
		width, height          = 10240, 10240
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func newton(z complex128) color.Color {
	const iterations = 1024
	const contrast = 30
	const threshold = 0.00001

	v := z
	for n := uint16(0); n < iterations; n++ {
		v = v*v*v*v - 1

		if r, theta := cmplx.Polar(v); r < threshold {
			val := uint8(255 - contrast*n)
			// Color according to quarter of root
			switch {
			case theta < -math.Pi/2.0:
				// Red
				return color.RGBA{val, 0, 0, 255}
			case theta < 0.0:
				// Green
				return color.RGBA{0, val, 0, 255}
			case theta < math.Pi/2.0:
				// Blue
				return color.RGBA{0, 0, val, 255}
			default:
				// Yellow
				return color.RGBA{val, val, 0, 255}

			}
		}
	}
	return color.Black
}
