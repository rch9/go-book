package popcount

var Pc [256]byte

func init() {
    for i := range Pc {
        Pc[i] = Pc[i/2] + byte(i&1)
    }
}
// IDEA: 
func PopCount(x uint64) (res int) {
    for i := 0; i < 8; i++ {
        res += int(Pc[byte(x>>(uint8(i)*8))])
    }
    return res
}

// func test(d int) int {
//     fmt.Printf("%b\n", byte(d))
//     fmt.Printf("%b\n", byte(d>>8))
//     fmt.Printf("%b\n", d)
//     return 5
// }
