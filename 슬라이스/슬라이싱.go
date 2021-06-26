package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:3] // [1]에서 [3-1]까지 선택

	fmt.Println(array)
	fmt.Println(slice)

}
