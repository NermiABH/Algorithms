package main

import (
	"fmt"
)

//Input - the number of the fibonacci number
//Output - the fibonacci number itself
//The algorithm is implemented on the basis of matrices

func main() {
	var (
		a, b, c = 0, 1, 1 // current Matrix
		d, e, f = 1, 2, 3 //shift Matrix
		num     int
	)
	fmt.Println("Введите число фиббоначи")
	fmt.Scan(&num)

	finish := num / 3
	for i := 1; i <= finish; i++ {
		a, b, c = a*d+b*e, a*e+b*f, b*e+c*f
	}
	if finish*3 == num {
		fmt.Println(a)
	} else if finish*3+1 == num {
		fmt.Println(b)
	} else {
		fmt.Println(c)
	}
}
