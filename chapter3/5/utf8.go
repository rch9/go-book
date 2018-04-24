package main

import "fmt"

func main() {
	// fmt.Println("\u4e16")
	// fmt.Println("\x11")
	// fmt.Println("\100")
	//
	// fmt.Println(HasPrefix("string", "st"))
	// fmt.Println(HasPrefix("string", "ing"))

	// s := "hello world,\u4141\u4122"
	// fmt.Println(len(s))                    // 3
	// fmt.Println(utf8.RuneCountInString(s)) // 1
	//
	// for i := 0; i < len(s); {
	// 	r, size := utf8.DecodeRuneInString(s[i:])
	// 	fmt.Printf("%d\t%c\n", i, r)
	// 	i += size
	// }
	//
	// for i, r := range s {
	// 	fmt.Printf("%d\t%q\t%d\n", i, r, r)
	// }

	s := "hello world"
	fmt.Printf("% x\n", s)
	r := []rune(s)
	fmt.Printf("%x\n", r)

	fmt.Println(string(1234567))
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
