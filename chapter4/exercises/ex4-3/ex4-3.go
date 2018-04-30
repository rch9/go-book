package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	reverse(&a)
	fmt.Println(a)
}

func reverse(s *[4]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
		// (*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
