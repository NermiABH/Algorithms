package main

import "fmt"

func BubbleSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func main() {
	slice := []int{2, 1, 20, 11, 6, 9, 3, 4, 8, 18, 15}
	fmt.Print("Not sorted slice:\n", slice)
	BubbleSort(slice)
	fmt.Print("\nSorted Array:\n", slice)
}
