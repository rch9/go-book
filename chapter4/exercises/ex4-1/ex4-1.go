package main

import (
	"crypto/sha256"
	"fmt"
)

func countDiffInBytes(a1, a2 byte) (res int) {
	for a1 != 0 {
		if a1&1 == a2&1 {
			res++
		}
		a1 >>= 1
		a2 >>= 1
	}
	return res
}

func countDiffInDigests(d1, d2 *[32]byte) (res int) {
	for i := 0; i < len(d1); i++ {
		res += countDiffInBytes(d1[i], d2[i])
	}
	return res
}

func main() {
	d1 := sha256.Sum256([]byte("x"))
	d2 := sha256.Sum256([]byte("X"))

	fmt.Println(countDiffInDigests(&d1, &d2))
}
