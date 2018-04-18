package main

import (
    "fmt"
    "os"
    "log"
)

var cwd string

func init() {
    var err error
    cwd, err = os.Getwd()
    if err != nil {
        log.Fatalf("error os.Getwd: %v", err)
    }
    fmt.Println(cwd)
}

func main() {
    // x := "hello!"
    // for i := 0; i < len(x); i++ {
    //     x := x[i]
    //     if x != '!' {
    //         x := x + 'A' - 'a'
    //         fmt.Printf("%c", x)
    //     }
    // }
}
