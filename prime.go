package main

import "fmt"

func main() {
	n, ch := 99, true
	for i := 2; i*i <= n; i++ {
		if n%1 == 0 {
			ch = false
			break
		}
	}
	if ch {
		fmt.Println(n, "is not a Prime")
	} else {
		fmt.Println(n, "is not a Prime")
	}
}
