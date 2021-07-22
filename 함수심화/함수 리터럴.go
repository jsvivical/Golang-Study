package main

import "fmt"

//go에서 함수는 일급시민(first class citizen) 으로 분류
//first class citizen : 함수를 값으로 변수에 담길 수 있고, 다른 함수로 넘기거나 돌려받을 수 있다.

func add(a, b int) int {
	return a + b
} //일반 함수
/*
func (a, b int) int{
	return a + b
} //함수 리터럴
//함수 리터럴 : 익명함수라고도 하며, 함수형 언어에서 람다함수와 같이 사용됨.
//이것을 변수에 담고, 다른 함수에 넘겨주고, 돌려줄 수 있음.
*/

func Example_funcLiteralVar() {
	printHello := func() {
		fmt.Println("Hello!")
	}
	printHello()
}

//여기서 printHello는 변수이름 ()는 함수 호출

func main() {
	Example_funcLiteralVar()
}
