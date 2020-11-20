// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
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

type plotFunc func(float64, float64) (float64, int, bool)

func getFunc(r *http.Request) plotFunc {
	result := foo

	if err := r.ParseForm(); err != nil {
		log.Print(err)
		return result
	}

	funcName := "default"

	for k, v := range r.Form {
		if k == "drawing" {
			funcName = v[0]
		}
	}

	switch funcName {
	case "default":
		result = foo
	case "eggbox":
		result = eggBox
	case "moguls":
		result = moguls
	case "saddle":
		result = saddle
	default:
		log.Printf("Error: Invalid plot choice %q.\n", funcName)
		return result
	}

	return result
}

func corners(i, j int, f plotFunc) (float64, float64, float64, float64, float64, float64, float64, float64, int, bool) {
	var ax, ay, bx, by, cx, cy, dx, dy float64
	var col int // color
	var ok bool

	if ax, ay, col, ok = corner(i+1, j, f); !ok {
		return 0, 0, 0, 0, 0, 0, 0, 0, 0, false
	}
	if bx, by, col, ok = corner(i, j, f); !ok {
		return 0, 0, 0, 0, 0, 0, 0, 0, 0, false
	}
	if cx, cy, col, ok = corner(i, j+1, f); !ok {
		return 0, 0, 0, 0, 0, 0, 0, 0, 0, false
	}
	if dx, dy, col, ok = corner(i+1, j+1, f); !ok {
		return 0, 0, 0, 0, 0, 0, 0, 0, 0, false
	}

	return ax, ay, bx, by, cx, cy, dx, dy, col, true
}

func svgHead() string {
	return fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
}

func svgPolygon(ax, ay, bx, by, cx, cy, dx, dy float64, col int) string {
	return fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke='#%06x'/>\n",
		ax, ay, bx, by, cx, cy, dx, dy, col)
}

func svgTail() string {
	return "</svg>"
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	f := getFunc(r)

	var err error

	if _, err = io.WriteString(w, svgHead()); err != nil {
		log.Print(err)
		return
	}

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			var ax, ay, bx, by, cx, cy, dx, dy float64
			var col int
			var ok bool

			if ax, ay, bx, by, cx, cy, dx, dy, col, ok = corners(i, j, f); !ok {
				continue
			}

			if _, err = io.WriteString(w, svgPolygon(ax, ay, bx, by, cx, cy, dx, dy, col)); err != nil {
				log.Print(err)
				return
			}
		}
	}

	if _, err = io.WriteString(w, svgTail()); err != nil {
		log.Print(err)
		return
	}
}

func corner(i, j int, f plotFunc) (float64, float64, int, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z & color col
	z, col, ok := f(x, y)
	if !ok {
		return 0, 0, 0, false
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, col, true
}

// Compute color intensity according to score in [0.0, 1.0] range
// Add power to make colors more distinct
func colorIntensity(score float64) int {
	const (
		power        = 3.0
		maxIntensity = 255
	)
	return int(math.Pow(score, power) * maxIntensity)
}

// Give z-value a color according to its "distance" from the extreme values and from the mid-value
func color(z, minz, maxz float64) int {
	const (
		redShift   = 16
		greenShift = 8
		blueShift  = 0
	)

	// Scores
	redScore := (z - minz) / (maxz - minz)                      // Distance from max
	greenScore := math.Abs(z-(maxz+minz)/2) / (maxz - minz) / 2 // Distance from mid
	blueScore := (maxz - z) / (maxz - minz)                     // Distance from min

	// Intensities
	redIntensity := colorIntensity(redScore)
	greenIntensity := colorIntensity(greenScore)
	blueIntensity := colorIntensity(blueScore)

	// Combine intensities to RGB value
	return redIntensity<<redShift +
		greenIntensity<<greenShift +
		blueIntensity<<blueShift
}

func foo(x, y float64) (float64, int, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	r = math.Sin(r) / r
	if math.IsNaN(r) || math.IsInf(r, 0) {
		return 0, 0, false
	}
	return r, color(r, -0.22, 0.99), true
}

func eggBox(x, y float64) (float64, int, bool) {
	r := .5 * math.Sin(.42*x) * math.Sin(.42*y)
	if math.IsNaN(r) || math.IsInf(r, 0) {
		return 0, 0, false
	}
	return r, color(r, -0.5, 0.5), true
}

func moguls(x, y float64) (float64, int, bool) {
	r := -math.Abs(.35 * math.Sin(.42*x) * math.Sin(.42*y))
	if math.IsNaN(r) || math.IsInf(r, 0) {
		return 0, 0, false
	}
	return r, color(r, -0.35, 0), true
}

func saddle(x, y float64) (float64, int, bool) {
	r := -math.Sin(.15*x) * math.Sin(.15*y)
	if math.IsNaN(r) || math.IsInf(r, 0) {
		return 0, 0, false
	}
	return r, color(r, -1., 1.), true
}
