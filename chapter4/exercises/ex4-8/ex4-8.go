package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

const (
	digit  = "digit"
	letter = "letter"
	space  = "space"
	punch  = "punch"
)

func main() {
	counts := make(map[string]int)  // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		switch true {
		case unicode.IsDigit(r):
			counts[digit]++
		case unicode.IsLetter(r):
			counts[letter]++
		case unicode.IsSpace(r):
			counts[space]++
		case unicode.IsPunct(r):
			counts[punch]++

		}
		utflen[n]++
	}
	fmt.Printf("type\t\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t\t%d\n", c, n)
	}
	fmt.Print("\nlen\t\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
