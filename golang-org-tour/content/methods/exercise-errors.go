// +build OMIT

package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
    var z, n float64 = 1, 0
	
	switch {
	case x == 0:
	    return 0, nil
	case x > 0:
	    for i := 0; i < 10; i++ {
            n = z-(z*z-x)/z/2
		    z = n
	    }
		return z, nil
	default:
	    return x, ErrNegativeSqrt(x)
	}	
    //return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
