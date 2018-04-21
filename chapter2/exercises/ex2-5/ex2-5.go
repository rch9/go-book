// х&(х-1)
package main

import (
    "fmt"
    "./popcount"
)

func main() {
    // popcount.PopCount()
    var x uint64 = 11
    fmt.Printf("\n%b, %b, %d\n", x&(x-1), x, popcount.PopCount(x))
}
