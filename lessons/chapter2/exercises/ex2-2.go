// $ go run ex2-2.go -i cmd -cu curr 3.1
// $ go run ex2-2.go -cu curr
// $ 1
// $ 10
//

package main

import (
	"./currconv"
	"./tempconv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"bufio"
)

// input direction
const (
	cmd = "cmd"
	std = "std"
)

// convertion  inits
const (
	curr = "curr"
	temp = "temp"
    all = "all"
)

var inpdir = flag.String("i", "std", "input direction")
var convun = flag.String("cu", "all", "convertion units")

func main() {
	flag.Parse()


    fmt.Println(handleInputDir(inpdir))

}

func parse(args []string) (res string) {
	for _, arg := range args {
		if val, err := strconv.ParseFloat(arg, 64); err == nil {
			res += convert(convun, val)
		} else {
			fmt.Fprint(os.Stderr, "Parsing problem\n")
			os.Exit(1)
		}
	}
	return res
}

func handleInputDir(inp *string) (res string) {
    switch *inp {
    case cmd:
		res = parse(flag.Args())
    case std:
		res = parse(scanFromStd())
    default:
        fmt.Fprint(os.Stderr, "bad -i\n")
        os.Exit(1)
    }
    return res
}

func scanFromStd() (res []string) {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		res = append(res, input.Text())
	}
	return res
}

func convert(convType *string, val float64) (res string) {
	switch *convType {
	case curr:
		res = wrapCurr(val)
	case temp:
		res = wrapTemp(val)
	case all:
		res = fmt.Sprintf("%s,\n%s\n", wrapCurr(val), wrapTemp(val))
    default:
        fmt.Fprint(os.Stderr, "bad -cu\b")
        os.Exit(1)
	}
    return res
}

func wrapCurr(val float64) string {
	return fmt.Sprintf("%s = %s,\n%s = %s",
		currconv.Btc(val), currconv.BTCtoUSD(currconv.Btc(val)),
		currconv.Usd(val), currconv.USDtoBTC(currconv.Usd(val)))
}

func wrapTemp(val float64) string {
	return fmt.Sprintf("%s = %s,\n%s = %s",
		tempconv.Kelvin(val), tempconv.KtoC(tempconv.Kelvin(val)),
		tempconv.Celsius(val), tempconv.CtoK(tempconv.Celsius(val)))
}
