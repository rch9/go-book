package main

import "fmt"

var m = make(map[string]int)

func main() {
	// str := []string{"a", "b", "c"}
	// fmt.Println(k(str))
	Add([]string{"a"})
	Add([]string{"a"})
	fmt.Println(Count([]string{"a"}))
	fmt.Println(k([]string{"a"}))
}

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}
