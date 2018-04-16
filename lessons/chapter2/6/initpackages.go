package main

import (
    "./popcount"
    "fmt"
)

// var a = b + c
// var b = f()
// var c = 1
//
// func f() int { return c + 1 }
//
// func init() {
//     fmt.Println("Hello")
// }
//
// func init() {
//     fmt.Println("world")
// }

func main() {
    fmt.Printf("\nres: %d\n", popcount.PopCount(1000))
    // var s byte = 500
    // fmt.Println(byte(s))
    // fmt.Println(popcount.Pc)
    // var s byte = 255
    // fmt.Printf("\n%b\n", s<<0)
    // fmt.Printf("\n%b\n", s<<1)
    // fmt.Printf("\n%b\n", s>>1)
}
