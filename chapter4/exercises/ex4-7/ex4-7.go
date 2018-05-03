package main

import "fmt"

//DOES NOT WORK!!!!

func main() {
	s := "123"
	fmt.Println(s)
	b := reverse([]byte(s))
	fmt.Printf("%s\n", b)
	//нужна память под byte, срез по string тоже string
}

func reverse(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}
