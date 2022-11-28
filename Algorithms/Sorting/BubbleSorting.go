package main

import "fmt"

func main() {
	var (
		slice      []int
		newElement int
	)
	for true {
		_, err := fmt.Scanln(&newElement)
		if err != nil {
			break
		}
		slice = append(slice, newElement)
	}

	for i := 0; i < len(slice)-1; i++ {
		for j := i; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	fmt.Println(slice)
}
