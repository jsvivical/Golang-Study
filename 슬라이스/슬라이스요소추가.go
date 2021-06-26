package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3}

	slice2 := append(slice, 4)
	slice2 = append(slice2, 5)
	slice2 = append(slice2, 6)

	for i, v := range slice2 {
		fmt.Printf("[%3d] %3d\n", i, v)
	}
}
