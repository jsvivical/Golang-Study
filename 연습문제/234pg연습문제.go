package main

import "fmt"

func main() {
	var i, j int = 5, 0

	for ; i > 0; i-- {
		for j = 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
