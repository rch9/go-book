package main

import "fmt"

func main() {
    fmt.Println(gcp(15, 6))
    fmt.Println(fib(3))
}

func gcp(x, y int) int {
    for y != 0{
        x, y = y, x%y
    }
    return x
}

func fib(n int) int {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        x, y = y, x+y
    }
    return x
}
