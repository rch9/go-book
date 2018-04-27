package main

import (
	"fmt"
	"time"
)

func main() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute

	fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration Sm0s"
	fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration lm0s"

	const (
		a = 1
		b
	)
	fmt.Println(a)
	fmt.Println(b)
}
