package main

import (
	"fmt"
	"math"
	// "os"
)

const (
	width, height = 600, 320 // размер изображения
	cells         = 100 // детализация сетки (кол-во ячеек)
	xyrange       = 30.0 // размер осей
	xyscale       = width / 2 / xyrange // масштаб изображения по xy
	zscale        = height * 0.4 // масштаб по z
	angle         = 30 * math.Pi / 180 //углы осей x, y. Что значит?
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := cornerAndZ(i+1, j)
			bx, by, bz := cornerAndZ(i, j)
			cx, cy, cz := cornerAndZ(i, j+1)
			dx, dy, dz := cornerAndZ(i+1, j+1)

			avg := getAvg(az, bz, cz, dz)
			norm := getNormByAvg(avg)
			color := getColorFromRangeByNormedAvg(norm)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%.6x' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Println("</svg>")
}

func cornerAndZ(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, у float64) float64 {
	r := math.Hypot(x, у) // r = sqrt(x^2 + y^2)

	if r < 10e-3 {
		r = 10e-3
	}

    return math.Sin(r) / r
}

func getAvg(az, bz, cz, dz float64) float64 {
	return (az + bz + cz +dz) / 4
}

const (
	minZ = -1.0
	maxZ = 1.0
)
func getNormByAvg(avg float64) float64 {
	return (avg - minZ) / (maxZ - minZ)
}

const (
	red = 0xff0000
	green = 0x00ff00
	blue = 0x0000ff
)
func getColorFromRangeByNormedAvg(norm float64) (color int) {
	rgbComp := 0xff
	var coeff int

	switch true {
	case norm < 0.25: //incr green
		color = blue
		coeff = 0x100
	case norm < 0.5: //dect blue
		color = green
		coeff = 0x1
		norm = 1 - norm
	case norm < 0.75: // incr red
		color = green
		coeff = 0x10000
	default: 		 // decr green
		color = red
		coeff = 0x100
		norm  = 1 - norm
	}
	color += coeff*int(math.Round(float64(rgbComp)*norm))
	return color
}
