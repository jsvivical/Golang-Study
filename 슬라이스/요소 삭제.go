package main

import "fmt"

func main() {
	//첫번째 방법
	slice := []int{1, 2, 3, 4, 5, 6}
	idx1 := 2
	idx2 := 3
	for i := idx1 + 1; i < len(slice); i++ {
		slice[i-1] = slice[i]
	}
	slice = slice[:len(slice)-1]
	fmt.Println(slice)

	//두번째방법 (append() 이용)
	slice = append(slice[:idx2], slice[idx2+1])
	fmt.Println(slice)
}
