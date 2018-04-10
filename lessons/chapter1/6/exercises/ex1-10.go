// ./fetchall https://coinmarketcap.com

package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main() {
    start := time.Now()
    ch := make(chan string)
    for _, url := range os.Args[1:] {
        go fetch(url, ch)
    }
    //почему это работает так (про копирующийся ch в fetch и цикл)
    for range os.Args[1:] {

        //что такое 0644?
        err := ioutil.WriteFile("file", []byte(<-ch), 0644)
        if err != nil {
            fmt.Sprintf("writing in file: %v", err)
        }

        // fmt.Println(<-ch)
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err)
        return
    }

    // что делает Discard (считать и проигнорировать?)
    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close()
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    // что делает %7d
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
