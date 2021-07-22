//함수타입 변수란 함수를 값으로 갖는 변수를 의미
//함수포인터라고도 부름

package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func getOperator(op string) func(int, int) int {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {
	var operator func(int, int) int
	operator = getOperator("*")

	var result = operator(3, 4)
	fmt.Println(result)
}
