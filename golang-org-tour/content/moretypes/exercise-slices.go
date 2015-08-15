// +build OMIT

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var s []uint8
	var r [][]uint8 = make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		s = make([]uint8, dx)
		r[y] = s
		for x := 0; x < dx; x++ {
			s[x] = uint8((x + y)/2)
		}
	}
	return r
}

func main() {
	pic.Show(Pic)
}
