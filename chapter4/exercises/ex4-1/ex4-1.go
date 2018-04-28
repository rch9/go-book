package main

import "fmt"

func countDiff(a1, a2 int) (res int) {
	for a1 != 0 {
		if a1&1 == a2&1 {
			res++
		}
		a1 >>= 1
		a2 >>= 1
	}
	return res
}

func main() {
	fmt.Printf("%b\n%b\n", 0xD, 0xB)
	fmt.Println(countDiff(0xD, 0xB))
}
