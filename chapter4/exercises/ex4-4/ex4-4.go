package main

import "fmt"

func main() {
	// Rotate s left by two positions.
	s := []int{0, 1, 2, 3, 4}

	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s)

	s2 := []int{0, 1, 2, 3, 4}
	s2 = rotate(s2, 2)
	fmt.Println(s2)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, r int) (s2 []int) {
	s2 = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		s2[(i+r+1)%(len(s))] = s[i]
	}
	return s2
}
