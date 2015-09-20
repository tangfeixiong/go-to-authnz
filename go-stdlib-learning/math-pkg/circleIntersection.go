package main

import (
    "math"
    "fmt"
)

// circle1: x * x + y * y + 6 * x - 16 = 0
// circle2: x * x + y * y - 4 * x - 5 = 0
func main() {
    x := pointXofCirclesIntersection()
    fmt.Println("x=", x)
    y1, y2 := pointYofCircle1Intersection(x)
    fmt.Println("y of circle1 intersection:", y1, y2)
    y1, y2 = pointYofCircle2Intersection(x)
    fmt.Println("y of circle2 intersection:", y1, y2)
}

func pointXofCirclesIntersection() float64 {
    return float64(16  - 5)/float64(6 - -4)
}

func pointYofCircle1Intersection(x float64) (float64, float64) {
    power := math.Pow(x+3, 2)
    y := math.Sqrt(25.0 - power)
    return y, -y
}

func pointYofCircle2Intersection(x float64) (float64, float64) {
    power := math.Pow(x-2, 2)
    y := math.Sqrt(9.0 - power)
    return y, -y
}