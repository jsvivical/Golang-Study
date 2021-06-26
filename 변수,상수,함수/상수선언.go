package main

import "fmt"

func main() {
	fmt.Println()

	const C int = 10

	var b int = C * 20
	fmt.Println(b, &C) //변수가 아닌 상수의 메모리에는 접근할 수 없기 때문에 출력하면 에러 발생
}
