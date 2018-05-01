package main

import "fmt"

type Point struct{ X, Y int }

type address struct {
	hostname string
	point    Point
}

func main() {
	p1 := Point{1, 2}
	p2 := Point{1, 2}
	fmt.Println(p1 == p2) // can use as a keys in maps

	a1 := address{"a", Point{1, 2}} // can cmp
	a2 := address{"a", Point{1, 2}}
	fmt.Println(a1 == a2)

	hits := make(map[address]int)
	hits[address{"a", Point{2, 2}}]++
	fmt.Println(hits)
}
