package main

import (
    "time"
    "fmt"
    "os"
    "strings"
)

func echo1(args []string) string {
    var s, sep string
    for i := 1; i < len(args); i++ {
        s += sep + args[i]
        sep = " "
    }
    return s
}

func echo2(args []string) string {
    s, sep := "", ""
    for _, arg := range args[1:] {
        s += sep + arg
        sep = " "
    }
    return s
}

func echo3(args []string) string {
    return strings.Join(args[1:], " ")
}

// 33.046193ms
// 33.403516ms
// 12.803524ms

func main() {
    var timeInterval time.Time
    var t1, t2, t3 time.Duration

    for i := 0; i < 100000; i++ {
        timeInterval = time.Now()
        echo1(os.Args)
        t1 += time.Since(timeInterval)

        timeInterval = time.Now()
        echo2(os.Args)
        t2 += time.Since(timeInterval)

        timeInterval = time.Now()
        echo3(os.Args)
        t3 += time.Since(timeInterval)
    }

    fmt.Println(t1)
    fmt.Println(t2)
    fmt.Println(t3)
}
