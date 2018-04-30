package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "h e   l  l  o  w    o    r   l    d   "

	b := []byte(str)
	fmt.Printf("%s\n", removeDubsSpaces(b))
}

func removeDubsSpaces(b []byte) []byte {
	for i := 0; i < len(b)-1; i++ {
		if b[i] == b[i+1] && unicode.IsSpace(rune(b[i])) {
			b = remove(b, i)
			i--
		}
	}
	return b
}

func remove(slice []byte, i int) []byte {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
