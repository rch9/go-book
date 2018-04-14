package main

import "fmt"

func newInt() *int {
    return new(int)
}

func main() {
    fmt.Println(newInt())
}
