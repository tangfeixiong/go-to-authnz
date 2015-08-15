// +build OMIT

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    sum := []int{0, 1}
    return func() int {
	v := sum[0] + sum[1]
	sum[0], sum[1] = sum[1], v
	return v
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Print(f(), ", ")
    }
    fmt.Println()
}
