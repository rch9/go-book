// Embed demonstrates basic struct embedding.
package main

import "fmt"

type point struct{ X, Y int }

type circle struct {
	point
	Radius int
}

type Wheel struct {
	circle
	Spokes int
}

func main() {
	var w Wheel
	//!+
	w = Wheel{circle{point{8, 8}, 5}, 20}

	w = Wheel{
		circle: circle{
			point:  point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
	}

	fmt.Printf("%#v\n", w)
	// Output:
	// Wheel{Circle:Circle{Point:Point{X:8, Y:8}, Radius:5}, Spokes:20}

	w.X = 42

	fmt.Printf("%#v\n", w)
	// Output:
	// Wheel{Circle:Circle{Point:Point{X:42, Y:8}, Radius:5}, Spokes:20}
	//!-
}
