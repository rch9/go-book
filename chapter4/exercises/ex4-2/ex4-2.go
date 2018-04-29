package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
	"strings"
)

var flagSha = flag.Int("sha", 256, "sha***")

const (
	_ = iota * 128
	_
	SHA256
	SHA384
	SHA512
)

func main() {
	flag.Parse()
	str := makeDataToSha(flag.Args(), " ")
	// fmt.Println(*flagSha, str)

	fmt.Println(sha(*flagSha, []byte(str)))
}

func makeDataToSha(strs []string, sep string) (res string) {
	return strings.Join(strs, sep)
}

func sha(arg int, data []byte) (res []byte) {

	switch arg {
	case SHA256:
		arr := sha256.Sum256(data)
		res = arr[:]
	case SHA384:
		arr := sha512.Sum384(data)
		res = arr[:]
	case SHA512:
		arr := sha512.Sum512(data)
		res = arr[:]
	default:
		fmt.Fprintf(os.Stderr, "bad arg %d in sha()", arg)
	}

	return res
}
