package main

import (
    "./tempconv"
    "fmt"
)

func main() {
    fmt.Println(tempconv.AbsZeroK)
    fmt.Println(tempconv.KtoC(10))
    fmt.Println(tempconv.CtoK(100))
}
