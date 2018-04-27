package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Satarday
)

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
)

const (
	Q = 1 << (10 * iota)
	KiB
	MiB
)

func main() {
	fmt.Println(FlagUp)
	fmt.Println(KiB)
	fmt.Println(MiB)
}
