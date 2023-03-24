package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var a, b, c int
	fmt.Fscan(in, &a, &b, &c)
	if c < 0 {
		fmt.Println("NO SOLUTION")
	} else if (a+b) == c*c && (2*a+b) == c*c {
		fmt.Println("MANY SOLUTIONS")
	}
}
