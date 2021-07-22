package main

import "fmt"

func sum(nums ...int) int {
	sum := 0

	fmt.Printf("nums 타입 : %T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 20))
	fmt.Println(sum(1, 20, 100, 500, 200, 300, 400, 100, 100000))
}
