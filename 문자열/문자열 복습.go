package main

import "fmt"

func main() {
	//문자열은 uint8타입의 연속이다.
	//문자열은 두 개의 워드로 된 데이터 구조체이다.
	//첫번째 워드는 뒤에 숨겨져 있는 배열을 가리키는 포인터이고,
	//두번째 워드는 문자열의 길이이다.

	a := "Hello, World!"

	for _, v := range a {
		fmt.Printf("%c\n", v)
	}
	//한글은 rune
	b := "안녕, 세계!"
	newb := []rune(b) //rune 슬라이스로 타입 변환

	for _, v := range newb {
		fmt.Printf("%c\n", v)
	}
}
