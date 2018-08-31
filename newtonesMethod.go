package main

import (
	"fmt"
)

var i = 0
var prev float64 = 0

func Sqrt(x float64) float64 {
	z := float64(1)
	return rec(x, z)
}

func rec(x float64, z float64) float64 {
    prev = z
	y := z - (z*z-x)/(2*z)
    i++
	if prev == 1 || prev > y {
		return rec(x, y)
	}
	return y
}

func main() {
	fmt.Println("loop count = ", i, "result = ", Sqrt(16))
}
