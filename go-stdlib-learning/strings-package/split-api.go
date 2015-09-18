// https://github.com/golang/go/blob/master/src/strings/strings.go
// https://github.com/golang/go/blob/master/src/net/url/url.go

// strings split api vs net url api's split function

package main

import (
    "strings"
    "fmt"
)

func main() {
    u := make([]string, 0, 2)
    u = strings.Split("http://example.com/index.html#name", "#")
    fmt.Printf("%q %q \n", u[0], u[1]) 
    
    u[0], u[1] = split("http://example.com/index.html#name", "#", true)
    fmt.Printf("%q %q \n", u[0], u[1])
} 

// Maybe s is of the form t c u.
// If so, return t, c u (or t, u if cutc == true).
// If not, return s, "".
func split(s string, c string, cutc bool) (string, string) {
	i := strings.Index(s, c)
	if i < 0 {
		return s, ""
	}
	if cutc {
		return s[:i], s[i+len(c):]
	}
	return s[:i], s[i:]
}
