package main

import "fmt"

func changeArray(array2 *[5]int) {
	array2[2] = 200
}

func changeSlice(slice2 []int) {
	slice2[2] = 200
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}
	var p *[5]int
	p = &array
	fmt.Println("array")
	for i, v := range array {
		fmt.Printf("[%3d] %3d\n", i, v)
	}
	changeArray(p)
	for i, v := range array {
		fmt.Printf("[%3d] %3d\n", i, v)
	}

	fmt.Println("slice")

	for i, v := range slice {
		fmt.Printf("[%3d] %3d\n", i, v)
	}
	changeSlice(slice)
	for i, v := range slice {
		fmt.Printf("[%3d] %3d\n", i, v)
	}
}
