// Package brat implements Mandelbrot with complex64
package brat

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/big"
)

// Draw function
func Draw(w io.Writer, xmin, ymin, xmax, ymax float64, width, height int) {
	var x *big.Rat = new(big.Rat).SetFloat64(0)
	var y *big.Rat = new(big.Rat).SetFloat64(0)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y.SetFloat64(float64(py)/float64(height)*(ymax-ymin) + ymin)
		for px := 0; px < width; px++ {
			x.SetFloat64(float64(px)/float64(width)*(xmax-xmin) + xmin)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(x, y))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

var temp1 *big.Rat = new(big.Rat)
var temp2 *big.Rat = new(big.Rat)
var temp3 *big.Rat = new(big.Rat)
var two *big.Rat = new(big.Rat).SetFloat64(2)
var four *big.Rat = new(big.Rat).SetFloat64(4)

func isGtTwo(x, y *big.Rat) bool {
	if x.Cmp(two) <= 0 && y.Cmp(two) <= 0 {
		return false
	}

	temp1.Mul(x, x)
	temp2.Mul(y, y)
	temp3.Add(temp1, temp2)

	return temp3.Cmp(four) > 0
}

func mandelbrot(x, y *big.Rat) color.Color {
	const iterations = 200
	const contrast = 15

	var vx, vy *big.Rat

	vx = new(big.Rat).SetFloat64(0)
	vy = new(big.Rat).SetFloat64(0)

	for n := uint8(0); n < iterations; n++ {
		// 						v = v*v + z
		// vx*vx-vy*vy
		temp1.Mul(vx, vx)
		temp2.Mul(vy, vy)
		temp2.Neg(temp2)
		// 2*vx*vy
		temp3.Mul(vx, vy)
		temp3.Mul(temp3, two)
		// vx
		vx.Add(temp1, temp2)
		vx.Add(vx, x) // 				+ z
		// vy
		vy.Add(temp3, y) // 			+ z

		if isGtTwo(vx, vy) {
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
