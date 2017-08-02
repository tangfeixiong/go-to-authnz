package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	// Parse + String preserve the original encoding.
	u, err := url.Parse("https://example.com/foo%2fbar")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Path)
	//fmt.Println(u.RawPath) // encoded path hint (Go 1.5 and later only; see EscapedPath method)
	fmt.Println(u.String())
}

