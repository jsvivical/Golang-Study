package main

import "fmt"

func main() {

	b := [2][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}

	for I, arr := range b {
		for i, v := range arr {
			fmt.Printf("a[%d][%d] = %3d \t", I, i, v)
		}
		fmt.Println()

	}

}
