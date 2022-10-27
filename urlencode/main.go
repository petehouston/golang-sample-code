package main

import (
	"fmt"
	"net/url"
)

func main() {
	params := url.Values{}
	params.Add("author", "Pete Houston")
	params.Add("title", "URL Encode Query String Parameters in Go")

	queryStr := params.Encode()
	fmt.Println("URL Encoded: ", queryStr)

	originalQueryString := "author=Pete Houston&title=URL Encode Query String Parameters in Go"
	fmt.Println("Original: ", originalQueryString)
	encodedQueryString := url.QueryEscape(originalQueryString)
	fmt.Println("URL Encoded: ", encodedQueryString)

}
