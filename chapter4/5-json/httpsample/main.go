package main

import (
	"fmt"
	"net/url"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

func main() {

	q := url.QueryEscape("repo:golang/go is:open json decoder")
	fmt.Println(q)
	//repo%3Agolang%2Fgo+is%3Aopen+json+decoder

	// resp, err := http.Get(IssuesURL + "?q=" + "is:open json decoder")
	// if err != nil {
	// 	fmt.Println("error")
	// }
	// // fmt.Println(resp)
	//
	// var result IssuesSearchResult
	// if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
	// 	resp.Body.Close()
	// }
	// resp.Body.Close()
	// fmt.Printf("%d\n", result.Items[0].Number)
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
