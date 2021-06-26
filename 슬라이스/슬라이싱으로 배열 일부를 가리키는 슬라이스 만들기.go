package main

import "fmt"

func main() {

	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("슬라이싱으로 배열 \" 중간\"에 요소들 가리키기")
	slice1 := array[1:5]
	fmt.Println(slice1)
	fmt.Println("슬라이싱으로 배열 \"처음\"부터 n -1 까지의 요소들 가리키기")
	slice2 := array[:3]
	fmt.Println(slice2)

	fmt.Println("슬라이싱으로 배열 n부터 \"끝\"까지의 요소들 가리키기")
	slice3 := array[5:]
	fmt.Println(slice3)

	fmt.Println("슬라이싱으로 전체 요소들 가리키기")
	slice4 := array[:]
	fmt.Println(slice4)

	fmt.Println("슬라이싱으로 1부터 (3 - 1)까지의 인덱스를 가리키고 4만큼의 cap을 가짐")
	slice5 := array[1:3:4]
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))

}
