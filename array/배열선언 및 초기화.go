package main

import "fmt"

func main() {
	var nums [5]int
	for i := 0; i < 5; i++ {
		fmt.Print(nums[i], "\t")
	}
	fmt.Println()
	fmt.Println()
	days := [3]string{"monday", "tuesday", "wednesday"}

	for i := 0; i < 3; i++ {
		fmt.Print(days[i] + "\t")
	}
	fmt.Println()
	fmt.Println()

	var temps [5]float64 = [5]float64{24.6, 26.7}
	for i := 0; i < 5; i++ {
		fmt.Print(temps[i], "\t")
	}
	fmt.Println()
	fmt.Println()

	var s = [5]int{1: 10, 3: 30}
	//요소의 개수는 5개, 1번 인덱스는 10, 3번 인덱스는 30
	for i := 0; i < 5; i++ {
		fmt.Print(s[i], "\t")
	}
	fmt.Println()
	fmt.Println()

	x := [...]int{10, 20, 30} //배열 요소 개수 생략
	for i := 0; i < 3; i++ {
		fmt.Print(x[i], "\t")
	}

}
