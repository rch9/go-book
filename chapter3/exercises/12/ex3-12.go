package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "stop"
	s2 := "post"

	fmt.Println(chackAnogramm(s1, s2))

	s1 = "stop"
	s2 = "ssop"
	fmt.Println(chackAnogramm(s1, s2))
}

func chackAnogramm(s1, s2 string) bool { //ASCII
	if len(s1) == len(s2) {
		for _, v := range s1 {
			if !strings.Contains(s2, string(v)) {
				return false
			}
		}
		return true
	}
	return false
}
