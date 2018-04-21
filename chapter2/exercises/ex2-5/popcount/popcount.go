package popcount

func PopCount(x uint64) (res int) {
    for ;x != x&(x-1); x = x&(x-1) {
        res++
    }
    return res
}
