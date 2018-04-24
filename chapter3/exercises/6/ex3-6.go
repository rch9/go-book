package main

import (
    "image"
    "image/color"
    "image/png"
    "math/cmplx"
    "os"
)

// how to make const param in func
func main() {
    const (
        xmin, ymin, xmax, ymax = -2, -2, +2, +2
        width, height          = 1024, 1024
    )

    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py := 0; py < height; py++ {
        y := float64(py) / height * (ymax - ymin) + ymin //?
        for px := 0; px < width; px++ {
            x := float64(px) / width *(xmax - xmin) + xmin
            z := complex(x, y)
            img.Set(px, py, mandelbrot(z)) // TODO: add sum of mandelbrot / 4
        }
    }
    png.Encode(os.Stdout, img) // should check errors
}

// func supersampling(colors [4]color.Color) (res color.Color) {
//     var r, g, b, a uint
//
//     for color := range colors {
//         res.R
//     }
//
//     return res{r, g, b, a}
// }

// TODO: check pointers
// TODO: check syntax
func sumColors(rgba1, rgba2 color.RGBA) color.RGBA {
    rgba1.R += rgba2.R
    rgba1.G += rgba2.G
    rgba1.B += rgba2.B
    rgba1.A += rgba2.A

    return rgba1
}

func mandelbrot(z complex128) color.Color {
    const iterations  = 200
    const contrast = 15

    var v complex128
    for n := uint8(0); n < iterations; n++ {
        v = v*v + z
        if cmplx.Abs(v) > 2 {
            return color.Gray{255 - contrast*n}
        }
    }
    return color.Black
}
