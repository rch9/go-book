// Movie prints Movies as JSON.
package main

import (
	"encoding/json"
	"fmt"
)

//!+
type Movie struct {
	Title string "111"
	Year  int    `json:"222"`
	Yer   int    "333"
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Yer: 1},
	{Title: "Cool Hand Luke", Year: 1967, Yer: 2},
	{Title: "Bullitt", Year: 1968, Yer: 3},
}

func main() {

	// data1, _ := json.Marshal(movies)
	// fmt.Printf("%s\n", data1)

	data2, _ := json.MarshalIndent(movies, "", "    ")
	fmt.Printf("%s\n", data2)

	var titles []struct{ Year int }
	json.Unmarshal(data2, &titles)
	fmt.Println(titles[1])

	// {
	// 	data, err := json.Marshal(movies)
	// 	if err != nil {
	// 		log.Fatalf("JSON marshaling failed: %s", err)
	// 	}
	// 	fmt.Printf("%s\n", data)
	// }

	// {
	// 	//!+Marshal Indent
	// 	data, err := json.MarshalIndent(movies, "", "    ")
	// 	if err != nil {
	// 		log.Fatalf("JSON marshaling failed: %s", err)
	// 	}
	// 	fmt.Printf("%s\n", data)
	// 	//!-MarshalIndent
	//
	// 	//!+Unmarshal
	// 	var titles []struct{ Title string }
	// 	if err := json.Unmarshal(data, &titles); err != nil {
	// 		log.Fatalf("JSON unmarshaling failed: %s", err)
	// 	}
	// 	fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
	// 	//!-Unmarshal
	// }
}
