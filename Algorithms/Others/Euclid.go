package main

import (
	"fmt"
)

func euclid(a int, b int) int {
	var r int
	for b > 0 {
		r = a % b
		a, b = b, r
	}
	return a
}

func main() {
	a, b := 11000000, 1105000
	fmt.Println(a % b)
	fmt.Printf("Euclid(%d, %d) = %d\n", a, b, euclid(a, b))
	fmt.Println(a / euclid(a, b))
}
