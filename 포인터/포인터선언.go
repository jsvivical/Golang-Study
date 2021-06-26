package main

import "fmt"

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
