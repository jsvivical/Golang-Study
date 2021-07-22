//함수 리터럴은 이름없는 함수로 함수명을 적지 않고 함수 타입 변숫값으로 대입되는 함숫값을 의미
//익명함수 또는 람다 함수라고 불림

package main

import "fmt"

type opFunc func(a, b int) int

func getOperator(op string) opFunc {
	if op == "*" {
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	fn := getOperator("*")

	result := fn(3, 5)
	fmt.Println(result)
}
