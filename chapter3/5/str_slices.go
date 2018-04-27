package main

import (
	"fmt"
	"go-book/chapter3/5/basename1"
	"go-book/chapter3/5/basename2"
	"go-book/chapter3/5/comma"
)

func main() {
	fmt.Println(basename1.Basename("a/b/c.go"))
	fmt.Println(basename2.Basename("a/b/c.go"))
	fmt.Println(comma.Comma("1234567"))
}
