package main

import (
    "math"
    "fmt"
)

// Parabola Function: y * y = 2 * p * x
func main() {
    p, x := 2.0, 1.0
    _, angleParabola := computeDerivativeAndAngle(p, x)
    fmt.Printf("parabola: %v ---- %v", angleParabola, 180.0 - angleParabola * 2.0)
    fmt.Println()
    x = 0.5
    _, angleParabola = computeDerivativeAndAngle(p, x)
    fmt.Printf("parabola: %v ---- %v    ", angleParabola, 180.0 - angleParabola * 2.0)
    angleFocus := computeFocusAngle(p, x)
    fmt.Println("focus: ", angleFocus)
    p, x = 1.0, 0.4
    _, angleParabola = computeDerivativeAndAngle(p, x)
    fmt.Printf("parabola: %v ---- %v    ", angleParabola, 180.0 - angleParabola * 2.0)
    angleFocus = computeFocusAngle(p, x)
    fmt.Println("focus: ", angleFocus)    
}

func computeDerivativeAndAngle(p, x float64) (float64, float64) {
    derivative := math.Sqrt(2.0 * p) / 2 / math.Sqrt(x)
    radianParabola := math.Atan(derivative)
    angleParabola := radianParabola * 180 / math.Pi
    return derivative, angleParabola
}

func computeFocusAngle(p, x float64) float64 {
    y := math.Sqrt(2.0 * p * x)
    radian := math.Atan2(y - 0.0, math.Abs(x - p / 2))
    angle := radian * 180 / math.Pi
    return angle
}