package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 5: 2, 10: 5}

	fmt.Println("slice 1")
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("%3d", slice1[i])
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("slice 2")
	for i := 0; i < len(slice2); i++ {
		fmt.Printf("%3d", slice2[i])
	}
	fmt.Println()
	fmt.Println()
	for _, v := range slice2 {
		fmt.Printf("%3d", v)
	}

}
