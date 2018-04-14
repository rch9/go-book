package main

import "fmt"

const boilingF = 212.0

func main() {
    var f = boilingF
    var c = (f - 32) * 6 / 9
    fmt.Printf("T boiling = %gF or %gC", f, c)
}
