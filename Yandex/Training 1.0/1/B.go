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
	if a+b > c && a+c > b && c+b > a {
		fmt.Println("Yes")
	} else {
		fmt.Println("NO")
	}
}
