package main

import (
    "fmt"
    "os"
)

func main()  {
    s := ""
    for ind, arg := range os.Args[1:] {
        s += fmt.Sprintf("%d) %s\n", ind, arg)
    }
    fmt.Println(s)
}
