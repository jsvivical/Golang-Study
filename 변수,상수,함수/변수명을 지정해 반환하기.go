package main

import "fmt"

func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}

	result = a / b
	success = true
	return
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	c, success = Divide(9, 0)
	fmt.Println(c, success)
}
