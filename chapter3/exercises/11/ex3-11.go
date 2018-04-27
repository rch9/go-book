package main

import (
	"fmt"

	"./comma"
)

func main() {
	fmt.Println(comma.Comma("12345.6789"))
	fmt.Println(comma.Comma("-1.2345678"))
	fmt.Println(comma.Comma("-1234.567"))
	fmt.Println(comma.Comma("-123.4"))
	fmt.Println(comma.Comma("123"))
	fmt.Println(comma.Comma("12"))
	fmt.Println(comma.Comma("-1"))
}
