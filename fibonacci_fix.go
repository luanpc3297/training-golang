package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		// fmt.Println(z)
	}
	return z
}

func main() {
	p := 9.0
	fmt.Println(Sqrt(p))
	fmt.Println(math.Sqrt(p))
}
