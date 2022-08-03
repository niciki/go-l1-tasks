package main

import "fmt"

func main() {
	x := 2
	y := 7
	fmt.Printf("Before: x = %d, y = %d\n", x, y)
	x, y = y, x
	fmt.Printf("After change 1: x = %d, y = %d\n", x, y)
	// method for languages where method 1 is unavavailable
	x = x + y
	y = x - y
	x = x - y
	fmt.Printf("After change 2: x = %d, y = %d\n", x, y)
}
