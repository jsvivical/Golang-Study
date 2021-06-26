package main

import "fmt"

func main() {
	str := "Hello 월드!"
	for i, v := range str {
		fmt.Printf("[%2d] 타입 : %T 값 : %5d 문자 : %c \n", i, v, v, v)
	}
}
