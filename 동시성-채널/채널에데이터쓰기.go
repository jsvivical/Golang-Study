package main

import (
	"fmt"
	"time"
)

/*****************************************************************************************************************************************
 * 채널 c에 x라는 값을 쓰고 싶다면 그냥 c <- x라고 쓰면 된다.
 * 이 때 x와 c의 타입은 같아야 함.
*****************************************************************************************************************************************/

//함수 매개변수를 채널로 선언하려고 chan키워드를 붙임. 그 뒤에는 반드시 채널의 타입(int)을 명시
//c <- x 문장으로 x라는 값을 c라는 채널로 쓴다고 명시
//그리고 close()로 채널을 닫음
func writeToChannel(c chan int, x int) {
	fmt.Println(x)
	c <- x
	close(c)
	fmt.Println(x) //실행할 수 없음
	//두 번째 println함수를 실행할 수 없는 이유
	//c라는 채널에 쓴 값을 아무도 읽지 않음 ->writeToChannel()함수의 실행은 c <- x라는 문장에서 막혀버리기 때문
	//따라서 time.Sleep()라는 문장이 끝나면 프로그램은 더 이상 writeChannel()을 기다리지 않고 종료해버린다.
}

func main() {
	//c라는 채널 변수를 정의
	//항상 채널에 타입을 지정해야 함.
	c := make(chan int)

	go writeToChannel(c, 10)
	time.Sleep(1 * time.Second)
}
