package main

import "fmt"

func main() {
	s := []int{0, 0, 0, 1, 1, 1, 1, 2, 3, 4, 4}
	fmt.Println(removeDups(s))
}

func removeDups(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			s = remove(s, i)
			i--
		}
	}
	return s
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
