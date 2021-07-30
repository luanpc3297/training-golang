package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64 = 99
	var i = math.Sqrt(n)
	if int(i)*int(i) == int(n) {
		fmt.Println("la so chinh phuong")
	} else {
		fmt.Println("khong la so chinh phuong.So chinh phuong gan nhat la", (int(i+0.5))*(int(i+0.5)))
	}
}
