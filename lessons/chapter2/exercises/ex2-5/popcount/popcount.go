package popcount

func PopCount(x uint64) (res int) {
    for x != 0 {
        if x&1 == 1 {
            res++
        }
        x>>=1
    }
    return res
}
