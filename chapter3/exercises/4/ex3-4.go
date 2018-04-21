package main

import (
    "./surface"
    "log"
    "net/http"
    "math"
    "strconv"
    "fmt"
    "os"
)

func main() {
    //?color=red&angle=30&angle=3
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "image/svg+xml")

        q := r.URL.Query()
        colorStr := q.Get("color")
        angleStr := q.Get("angle")        
        fmt.Print(angleStr)


        //WHY?
        // parsing "": invalid syntax
        ang, err := strconv.ParseFloat(angleStr, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "can not conver angle to float, %v", err)
        }

        surface.Draw(w, colorStr, ang * math.Pi / 180)
    })
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
