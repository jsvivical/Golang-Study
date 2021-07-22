package main

import "fmt"

// 포인터는 오직 한가지 목적을 가지고 있다.
//공유, 프로그램의 경계를 가로질러 값을 공유하는 것이다.

func main() {
	var a int = 500
	var p *int

	if p == nil {
		fmt.Println("아직 포인터 p에 주소가 할당되지않았습니다.")
	}

	p = &a

	*p = 200
	fmt.Println(a)
}
