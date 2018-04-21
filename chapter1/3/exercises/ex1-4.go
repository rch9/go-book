package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    fileNames := make(map[string]string)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts, fileNames)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            fmt.Println(f.Name())
            if err != nil {
                fmt.Fprintf(os.Stderr, "dub2: %v\n", err)
                continue
            }
            countLines(f, counts, fileNames)
            f.Close()
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\t%s\n", n, line, fileNames[line])
        }
    }

}

func countLines(f *os.File, counts map[string]int,
    fileNames map[string]string) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
        if counts[input.Text()] > 1 {
            fileNames[input.Text()] = f.Name()
        }
    }
}

// func countLines(f *os.File, counts map[string]int, filename string, )  {
//
// }
