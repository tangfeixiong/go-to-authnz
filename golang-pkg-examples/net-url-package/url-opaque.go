package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	// Sending a literal '%' in an HTTP request's Path
	req := &http.Request{
		Method: "GET",
		Host:   "example.com", // takes precedence over URL.Host
		URL: &url.URL{
			Host:   "ignored",
			Scheme: "https",
			Opaque: "/%2f/",
		},
		Header: http.Header{
			"User-Agent": {"godoc-example/0.1"},
		},
	}
	out, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.Replace(string(out), "\r", "", -1))
}

