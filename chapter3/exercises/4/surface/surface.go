package surface

import (
	"fmt"
	"math"
	"io"
)

const (
	width, height = 600, 320
	cells         = 100 // детализация сетки (кол-во ячеек)
	xyrange       = 30.0 // размер осей
	xyscale       = width / 2 / xyrange // масштаб изображения по xy
	zscale        = height * 0.4 // масштаб по z
)

var angle float64
var sin30, cos30 float64

//Разные fmt.Fprintf, неверное кол-во параметров в каждом
// но в сумме верное - работает (1 stream)

func Draw(out io.Writer, color string, pangle float64) {

	angle = pangle
	sin30, cos30 = math.Sin(angle), math.Cos(angle)

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: %s; stroke-width: 0.7' "+
		"width='%d' height='%d'>", color)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprint(out, "</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f(x, y)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, у float64) float64 {
	r := math.Hypot(x, у) // r = sqrt(x^2 + y^2)
    return math.Sin(r) / r
}
