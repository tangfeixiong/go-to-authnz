package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func(mr MyReader) Read(b []byte) (int, error) {
    if b != nil {
	    for i := 0; i < len(b); i++ {
		    b[i] = 0x41
		}
		return len(b), nil
	}
	return 0, nil
}

func main() {
	reader.Validate(MyReader{})
}
