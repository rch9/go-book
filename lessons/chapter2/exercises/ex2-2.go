package main

import (
    // "./currconv"
    // "./tempconv"
    "fmt"
    // "os"
    "flag"
    // "strings"
)


const curr = "curr"
const temp = "temp"

// QUESTION: fmt.Println(flag.Args()) ??

var choice = flag.String("ch", "111", "smth")

func main() {
    flag.Parse()

    fmt.Println(flag.Args())
    // if *choice {
    //     switch flag.Args() {
    //     case curr:
    //         fmt.Println(curr)
    //     case temp:
    //         fmt.Println(temp)
    //     }
    // }
}
