package main

import "fmt"

func main() {
	var slice []int

	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}

	slice[1] = 10 //없는 인덱스에 접근 -> 오류발생
	fmt.Println(slice)
}
