package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var (
		a, b int
		c    string
	)
	fmt.Fscan(in, &a, &b, &c)
	switch c {
	case "heat":
		fmt.Println(math.Max(float64(a), float64(b)))
	case "freeze":
		fmt.Println(math.Min(float64(a), float64(b)))
	case "auto":
		fmt.Print(b)
	case "fan":
		fmt.Println(a)
	}

}
