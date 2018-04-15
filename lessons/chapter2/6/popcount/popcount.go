package popcount

var Pc [256]byte

func init() {
    for i := range Pc {
        Pc[i] = Pc[i/2] + byte(i&1)
    }
}

func popCount(x uint64) int {
    return int(pc)
}
