package main

import "fmt"

func main() {
	// months := [...]string{1: "Jan", 2: "Fer", 3: "Mar", 4: "Apr", 5: "May"}
	//
	// fmt.Println(months[:6])
	// winter := months[:3]
	// winterAndSpring := winter[:6]
	// fmt.Println(winterAndSpring)

	s := []int{1, 2, 3, 4}
	a := [4]int{1, 2, 3, 4}
	reverse(s)
	fmt.Println(s)
	reverse2(a)
	fmt.Println(a)
	reverse(a[:])
	fmt.Println(a)

	// make([]T, len)
	// make([]T, cap, len)
	// make([]T, cap)[:len]
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse2(s [4]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
