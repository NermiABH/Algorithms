package main

import "fmt"

func BinarySearchNoRec(slice []int, elem int) int {
	first, last := 0, len(slice)-1
	middle := 0
	for first <= last {
		middle = (first + last) / 2
		if elem > slice[middle] {
			first = middle + 1
		} else if elem < slice[middle] {
			last = middle - 1
		} else {
			return middle
		}
	}
	return -1
}

func main() {
	slice := []int{1, 2, 3, 4, 6, 8, 9, 11, 15, 18, 20}
	fmt.Println(BinarySearchNoRec(slice, 6))
}
