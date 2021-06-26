package main

import "fmt"

//rune 타입 : 문자 하나를 표현하는데 사용
//rune타입은 4바이트 정수 타입인 int32의 별칭 타입(즉, 같은 타입)

func main() {
	var char rune = '한'

	fmt.Printf("%T\n", char) //%T 는 타입이름을 출력
	fmt.Println(char)
	fmt.Printf("%c\n", char)

}
