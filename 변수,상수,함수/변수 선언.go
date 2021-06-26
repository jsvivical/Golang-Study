package main

import "fmt"

func main() {
	//go 언어에서 변수명을 지을 때의 규칙
	//변수명은 문자, _ , 숫자를 사용해 지을 수 있지만, 첫 글자는 반드시 문자나 _로 시작
	//_ 를 제외한 어떤 특수문자도 불가능
	//왠만하면 영어로 사용
	//두번째 단어부터는 대문자로 사용(자바처럼)
	//되도록 _은 사용하지 않음

	var a int = 10
	var msg string = "hello variables"

	a = 20
	msg = "good morning"

	fmt.Print(msg, a)
}
