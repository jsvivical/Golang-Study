package main

import "fmt"

var cnt int = 3

func IncreaseAndReturn() int {
	fmt.Print()
	fmt.Println("increaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {
	if false && IncreaseAndReturn() < 5 {
		fmt.Println("1증가")
	}
	fmt.Println(cnt)
}
