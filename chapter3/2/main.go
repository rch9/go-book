package main

import "fmt"

func main() {
    var f float32 = 16777216
    fmt.Println(f)
    fmt.Println(f+1)
    fmt.Println(f == f+1)
    w := 0.3e10
    fmt.Printf("\n%0.5g, %0.5[1]e, %-1.5[1]f", w)
}
