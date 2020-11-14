// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type plotFunc func(float64, float64) (float64, bool)

func getFunc() plotFunc {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Error: Expecting 1 argument.")
		os.Exit(1)
	}

	var f plotFunc

	switch os.Args[1] {
	case "default":
		f = foo
	case "eggbox":
		f = eggBox
	case "moguls":
		f = moguls
	case "saddle":
		f = saddle
	default:
		fmt.Fprintf(os.Stderr, "Error: Invalid plot choice %q.\n", os.Args[1])
		os.Exit(2)
	}

	return f
}

func corners(i, j int, f plotFunc) (float64, float64, float64, float64, float64, float64, float64, float64, bool) {
	var ax, ay, bx, by, cx, cy, dx, dy float64
	var ok bool

	if ax, ay, ok = corner(i+1, j, f); !ok {
		return 0, 0, 0, 0, 0, 0, 0, 0, false
	}
	if bx, by, ok = corner(i, j, f); !ok {
		return 0, 0, 0, 0, 0, 0, 0, 0, false
	}
	if cx, cy, ok = corner(i, j+1, f); !ok {
		return 0, 0, 0, 0, 0, 0, 0, 0, false
	}
	if dx, dy, ok = corner(i+1, j+1, f); !ok {
		return 0, 0, 0, 0, 0, 0, 0, 0, false
	}

	return ax, ay, bx, by, cx, cy, dx, dy, true
}

func svgHead() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
}

func svgPolygon(ax, ay, bx, by, cx, cy, dx, dy float64) {
	fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
		ax, ay, bx, by, cx, cy, dx, dy)
}

func svgTail() {
	fmt.Println("</svg>")
}

func main() {
	f := getFunc()

	svgHead()

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			var ax, ay, bx, by, cx, cy, dx, dy float64
			var ok bool

			if ax, ay, bx, by, cx, cy, dx, dy, ok = corners(i, j, f); !ok {
				continue
			}

			svgPolygon(ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	svgTail()

}

func corner(i, j int, f plotFunc) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, ok := f(x, y)
	if !ok {
		return 0, 0, false
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func foo(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	r = math.Sin(r) / r
	if math.IsNaN(r) || math.IsInf(r, 0) {
		return 0, false
	}
	return r, true
}

func eggBox(x, y float64) (float64, bool) {
	r := .5 * math.Sin(.42*x) * math.Sin(.42*y)
	if math.IsNaN(r) || math.IsInf(r, 0) {
		return 0, false
	}
	return r, true
}

func moguls(x, y float64) (float64, bool) {
	r := -math.Abs(.35 * math.Sin(.42*x) * math.Sin(.42*y))
	if math.IsNaN(r) || math.IsInf(r, 0) {
		return 0, false
	}
	return r, true
}

func saddle(x, y float64) (float64, bool) {
	r := -math.Sin(.15*x) * math.Sin(.15*y)
	if math.IsNaN(r) || math.IsInf(r, 0) {
		return 0, false
	}
	return r, true
}
