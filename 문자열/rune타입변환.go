package main

import "fmt"

func main() {
	str := "Hello World"

	runes := []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}
	//[]rune 타입은 string으로 변환가능한 타입임!
	fmt.Println(str)
	fmt.Println(string(runes))

	str2 := "[]rune Type is the type that can be converted to string "
	runes2 := []rune(str2)

	fmt.Println(str2)
	fmt.Println(string(runes2))

}
