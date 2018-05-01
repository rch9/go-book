package main

import "fmt"

type Point struct{ X, y int } // can not export y, can export X

func main() {

	p := Point{1, 2}  // only in this package
	p2 := Point{y: 1} // only in this package

	fmt.Println(p, p2)

	pp := &Point{1, 2}
	fmt.Println(pp)
	pp = new(Point)
	fmt.Println(pp)
}
