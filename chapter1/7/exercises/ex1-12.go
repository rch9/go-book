package main

import (
    // "fmt"
    "lissajous"
    "log"
    "net/http"
    "strconv"
    "os"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        // should check errors
        cycles, err := strconv.Atoi(r.URL.Query()["cycles"][0])

        if err != nil {
            os.Exit(13)
        }

        lissajous.Lissajous(w, float64(cycles))
    })
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
