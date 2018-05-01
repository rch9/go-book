package main

import "fmt"

type Point struct {
	X, Y int
}

type Center Point

type Circle struct {
	Point
	Radius int
}

func main() {
	c := Circle{Point{1, 2}, 3}
	c.Point.X = 1
	c.Y = 2
	c.Radius = 3
	fmt.Println(c)
	// d := Circle{Point.X: 1, Y: 2, Radius: 0} error
	// fmt.Println(d)

	center := Center{}
	fmt.Println(center)
}
