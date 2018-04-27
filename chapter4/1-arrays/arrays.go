package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBR
	RUR
	RUB = 44
)

func main() {
	// var a [3]int
	// var b [3]int = [3]int{1, 2}
	// c := [...]int{1, 2, 3}
	//
	// for i, v := range a {
	// 	fmt.Print(i)
	// 	fmt.Println(v)
	// }

	symbol := [...]string{USD: "$", EUR: "E", GBR: "G", RUR: "R", RUB: "T"}
	fmt.Println(RUB, symbol[RUB], symbol)

}
