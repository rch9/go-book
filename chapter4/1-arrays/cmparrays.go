package main

import "fmt"

func main() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{2, 1}
	// d := [3]int{1, 2}

	fmt.Println(a == b, a == c)
}
