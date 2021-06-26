package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var f float64 = 23451312412.1232

	fmt.Print("a : ", a, "b : ", b, "f : ", f)
	fmt.Println("a : ", a, "b : ", b, "f : ", f)
	fmt.Printf("a : %d b : %d, f : %.2f\n", a, b, f)
	/*서식문자
	%v : 데이터 타입에 맞춰서 기본 형태로 출력함
	%T : 데이터 타입을 출력
	%t : 불리언을 true/false로 출력
	%d : 십진수 출력
	%b : 이진수 출력
	%c : 유니코드 문자 출력
	%o : 8진수로 값을 출력
	%O
	%x : 16진수로 값을 출력
	%X
	%e : 지수 형태로 실숫값 출력
	%E
	%f : 실수형태로 실숫값 출력
	%F
	%g : 값이 큰 실수값은 지수형태로 출력, 작은 실수값은 실수값 그대로
	%G
	%s : 문자열 출력
	%q : 특수 문자 기능이 동작하지 않고 그대로 출력됨
	%p : 메모리 주솟값 출력

	제어코드
	\n : 줄바꿈
	\t : 탭을 삽입
	\\ : \자체를 출력
	\" : "출력


	*/
}
