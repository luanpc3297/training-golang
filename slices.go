package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("SL1:", slice1)
	slice2 := slice1[2:5]
	slice2[1] = 10
	fmt.Println("SL2:", slice2[6])
	fmt.Println("SL1:", slice1)
}
