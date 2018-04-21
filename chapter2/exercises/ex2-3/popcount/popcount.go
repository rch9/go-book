package popcount

var Pc [256]byte

func init() {
    for i := range Pc {
        Pc[i] = Pc[i/2] + byte(i&1)
    }
}

func PopCount(x uint64) (res int) {
    for i := 0; i < 8; i++ {
        res += int(Pc[byte(x>>(uint8(i)*8))])
    }
    return res
}
