package main

import (
	"fmt"
)

//The number of rows of the triangle is fed to the input,
//and the triangle itself is fed to the output from *

func main() {
	var row int
	fmt.Scan(&row)
	startColumn := 0
	finishColumn := 0
	for i := 1; i <= row+1; i++ {
		startColumn, finishColumn = row-i+1, row+i-1
		for j := 1; j < finishColumn; j++ {
			if j <= startColumn {
				fmt.Print(" ")
				continue
			}
			fmt.Print("*")
		}
		fmt.Println()
	}
}
