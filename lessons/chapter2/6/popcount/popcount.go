package popcount

import "fmt"

var Pc [256]byte

func init() {
    for i := range Pc {
        Pc[i] = Pc[i/2] + byte(i&1)
    }
    fmt.Println()
    fmt.Println(Pc)
    fmt.Println()
}

func PopCount(x uint64) int {
    fmt.Printf("\nX:\n%b\n", x)

    fmt.Print("B:")
    fmt.Printf("\n%b", Pc[byte(x>>(0*8))])
    fmt.Printf("\n%b", Pc[byte(x>>(1*8))])
    fmt.Printf("\n%b", Pc[byte(x>>(2*8))])
    fmt.Printf("\n%b", Pc[byte(x>>(3*8))])
    fmt.Printf("\n%b", Pc[byte(x>>(4*8))])
    fmt.Printf("\n%b", Pc[byte(x>>(5*8))])
    fmt.Printf("\n%b", Pc[byte(x>>(6*8))])
    fmt.Printf("\n%b", Pc[byte(x>>(7*8))])
    fmt.Println()
    fmt.Println(byte(x))




    return int(
        Pc[byte(x>>(0*8))] +
        Pc[byte(x>>(1*8))] +
        Pc[byte(x>>(2*8))] +
        Pc[byte(x>>(3*8))] +
        Pc[byte(x>>(4*8))] +
        Pc[byte(x>>(5*8))] +
        Pc[byte(x>>(6*8))] +
        Pc[byte(x>>(7*8))])
}
