package main

import "fmt"

// NOTE: IDKx

// kilobyte KB 1000 byte
// megabyte MB 1000 KB
// gigabyte GB 1000 MB
// terabyte TB 1000 GB

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	PiB
	EiB
	ZiB
	YiB
)

const (
	_ = 1 + (iota * iota)
	KB
	MB
	GB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println(MB)
}
