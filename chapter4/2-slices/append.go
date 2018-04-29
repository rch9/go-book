package main

func main() {
	// var runes []rune
	// for _, r := range "Hello" {
	// 	runes = append(runes, r)
	// }
	// fmt.Printf("%q\n", runes)

	// var x, y []int
	// for i := 0; i < 10; i++ {
	// 	y = appendInt(x, i)
	// 	fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
	// 	x = y
	// }

	// var a []int
	// a = append(a, 1)
	// a = append(a, 2, 3)
	// a = append(a, a...)
	//
	// fmt.Println(a)

	// var x []int
	// x = appendInt(x, 1, 2, 3)
	// x = appendInt(x, x...)
	// fmt.Println(x)
}

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	// z[len(x)] = y
	copy(z[len(x):], y)
	return z
}
