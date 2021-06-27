package main

import "fmt"

func main() {
	//첫번째 방법 (반복문 사용)
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))
	for i, v := range slice1 {
		slice2[i] = v
	}

	slice1[1] = 100

	fmt.Println(slice1)
	fmt.Println(slice2)

	//두번째 방법(append() 함수 이용)
	slice3 := append([]int{}, slice1[0], slice1[1], slice1[2], slice1[3], slice1[4])

	fmt.Println(slice3)

	//세번째 방법(copy() 함수 이용)
	var lengthOfCopiedSlice int
	slice4 := make([]int, len(slice1))
	lengthOfCopiedSlice = copy(slice4, slice1) //slice4에 slice1 복사, 반환값은 복사된 요소의 개수

	fmt.Println(slice4, lengthOfCopiedSlice)

}
