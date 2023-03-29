package main

import "fmt"

func SelectionSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		pos, min := i, slice[i]
		for j := i + 1; j < len(slice); j++ {
			if min > slice[j] {
				pos, min = j, slice[j]
			}
		}
		slice[i], slice[pos] = slice[pos], slice[i]
	}
}

func main() {
	slice := []int{2, 1, 20, 11, 6, 9, 3, 4, 8, 18, 15}
	fmt.Print("Not sorted slice:\n", slice)
	SelectionSort(slice)
	fmt.Print("\nSorted Array:\n", slice)
}
