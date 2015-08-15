// +build OMIT

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (mr *rot13Reader) Read(p []byte) (int, error) {
    var b []byte = make([]byte, len(p))

    n, err := mr.r.Read(b)
	if err != io.EOF {
	    for i := 0; i < n; i++ {
		    switch {
			case b[i] < (0x41 + 13) :
			    p[i] = b[i] + 13
			case b[i] < 0x61 :
			    p[i] = b[i] - 13
			case b[i] < (0x61 + 13) :
			    p[i] = b[i] + 13
			default:
			    p[i] = b[i] - 13
			}
	    }
	} 
	
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
