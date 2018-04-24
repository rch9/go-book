package main

import "fmt"

func main() {
	// len - number of bytes
	s1 := "hello world@,,,,,,"
	s2 := "hello"
	fmt.Println(len(s1))
	fmt.Println(len(s2))
	fmt.Println(s1[10])
	fmt.Println(s1[2:])
	fmt.Println("hey + " + s1[:2])

	s := "left"
	t := s
	s += ", right"
	fmt.Println(s)
	fmt.Println(t)

	// s[0] = '1'
}
