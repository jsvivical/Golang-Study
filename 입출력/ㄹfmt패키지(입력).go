// +build ignore
package main

import "fmt"

func main() {
	fmt.Println()
	var a int
	var b int

	fmt.Println("정수 두 개 입력")
	n, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(n, err) //err은 초기에 nil상태이지만
	} else {
		fmt.Println(n, err, a, b)
	}
	//n은 함수의 반환값으로 성공적으로 입력했을 때, 그 입력한 값 개수 반환
	//입력이 실패하면 에러를 반환
}
