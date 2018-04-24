package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

// FIXME: THERE ARE MANY ERRORS...

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py += 2 {
		y0 := float64(py+0)/height*(ymax-ymin) + ymin
		y1 := float64(py+1)/height*(ymax-ymin) + ymin
		// y2 := float64(py+2)/height*(ymax-ymin) + ymin
		// y3 := float64(py+3)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px += 2 {
			x0 := float64(px+0)/width*(xmax-xmin) + xmin
			x1 := float64(px+1)/width*(xmax-xmin) + xmin
			// x2 := float64(px+2)/width*(xmax-xmin) + xmin
			// x3 := float64(px+3)/width*(xmax-xmin) + xmin

			z0 := complex(x0, y0)
			z1 := complex(x0, y1)
			z2 := complex(x1, y0)
			z3 := complex(x1, y1)
			img.Set(px, py, calcMidColor([]color.Color{
				mandelbrot(z0), mandelbrot(z1), mandelbrot(z2), mandelbrot(z3)}))
		}
	}
	png.Encode(os.Stdout, img) // should check errors
}

func calcMidColor(colors []color.Color) color.Color {
	var r, g, b, a uint32
	for _, c := range colors {
		cr, cg, cb, ca := c.RGBA()
		r, g, b, a = r+cr, g+cg, b+cb, a+ca
	}
	len := uint32(len(colors))
	r, g, b, a = r/len, g/len, b/len, a/len

	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			//return color.YCbCr{255 - contrast*n, contrast * n, 255 - contrast*n}
			return color.RGBA{255 - contrast*n, 0, 255 - contrast*n, 255}
		}
	}
	return color.Black
}
