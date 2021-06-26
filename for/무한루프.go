package main

import "fmt"

func main() {
	var a int = 0

	for true {
		fmt.Print("loop\n")
		a++
		if a == 100 {
			fmt.Println("loopEnd")
			break
		}
	}
}
