// +build OMIT

package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
    var z, n float64 = 1, 0
	
    for i := 0; i < 10; i++ {
        n = z-(z*z-x)/z/2
	z = n
    }
    fmt.Println(z)
	
    for {
	n = z-(z*z-x)/z/2
	if ( z > n && z - n < 0.01 ) { return n }
	if ( n > z && n - z < 0.01 ) { return n }
	z = n
    }
}

func main() {
	fmt.Println(Sqrt(2))
}
