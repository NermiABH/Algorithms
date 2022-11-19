package main

import "fmt"

//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
//The sum of these multiples is 23.
//Find the sum of all the multiples of 3 or 5 below 1000.

var finite = 999

func MultipleOf(n int) int {
	countOfdiv := finite / n
	sumGauss := countOfdiv * (countOfdiv + 1) / 2
	return n * sumGauss
}

func main() {
	sumMultiples := MultipleOf(3) + MultipleOf(5) - MultipleOf(15)

	fmt.Println(sumMultiples)
}
