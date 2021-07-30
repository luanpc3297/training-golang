package main

import "fmt"

func main() {
	x1, x2, n := 1, 1, 20
	for ; n > 0; n-- {
		fmt.Println(x2)
		x1, x2 = x2, x1+x2
	}
	fmt.Println(x2)
}
