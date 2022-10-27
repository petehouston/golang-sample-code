package main

import "fmt"

func main() {
	mapArticle := make(map[string]string)

	mapArticle["author"] = "Pete Houston"

	author, ok := mapArticle["author"]
	fmt.Printf("Author: %s | ok: %t\n", author, ok)

	title, ok := mapArticle["title"]
	if ok {
		fmt.Printf("Title: %s | ok: %t\n", title, ok)
	} else {
		fmt.Printf("Key 'title' doesn't exist")
	}

}
