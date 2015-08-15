// +build OMIT

package main

import (
	"golang.org/x/tour/wc"
        "strings"
)

func WordCount(s string) map[string]int {
    var result map[string]int = make(map[string]int)
    var f = strings.Fields(s)
	var (
	    test bool
		value int
	)
	for i := 0; i < len(f); i++ {
	    value, test = result[f[i]]
		if test { 
		    result[f[i]] = value + 1 
		} else { 
		    result[f[i]] = 1 
		}
	}
	return result

	//return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
