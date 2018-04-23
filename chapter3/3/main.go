package main

import (
    "fmt"
    "math/cmplx"
)

func main() {
    var x complex128 = complex(1, 2)
    var y complex128 = complex(3, 4)
    fmt.Println(x*y)
    fmt.Println(imag(x*y))
    fmt.Println(real(x*y))

    z := 1i
    // m := 8
    fmt.Printf("%T\n", z+8)

    fmt.Println(cmplx.Sqrt(-1))
}
