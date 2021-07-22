package main

import (
	"io"
	"os" //os패키지는 프로그램에서 커맨드라인 인수를 읽고, strout에 접근하는 데 사용한다.
)

func main() {
	myString := "" //회면에 출력할 텍스트를 담음
	//이 변수의 값은 프로그램의 커맨드라인 인수로 설정하거나 커맨드라인 인수가 없으면 코드 안에 적어둔 메시지로 설정
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument"
	} else {
		myString = arguments[1]
	}
	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}
