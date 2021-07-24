/********************************************************************************************************************
 * 파이프라인은 고루틴과 채널을 연결하는 가상 메서드로, 채널을 통해 한쪽 고루틴에서 출력된 다른 고루틴에 입력 가능
 * 파이프라인을 사용하면 채널을 실행할 때 상대방을 기다릴 필요가 없기 때문에 데이터흐름을 만들 수 있다
 * 주고 받는 값을 일일이 변수에 저장할 필요가 없어 변수 사용횟수를 줄일 수 있고, 메모리 절약가능
 * 설계가 간결해서 유지보수에 용이
 *아래의 프로그램은 주어진 범위에서 난수를 생성하다가 같은 값이 두번 나오면 멈추고, 그 동안 생성된 난수를 모두 더한 결과를 출력하고 종료
 *데이터는 파이프라인으로 구성한 채널을 통해 흐름
 *이 프로그램은 두 개의 채널을 사용함.
 * 첫 번째 함수는 첫번째 채널에서 난수를 가져와 두 번째 함수로 보냄.
 * 두번째 함수는 두 번째 채널에서 난수를 가져와서 세번째 함수로 보낸다.
 * 세번째 함수는 두번째 채널에서 값을 읽고, 숫자를 더하고,결과를 출력함

 ********************************************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// second() 함수가 first()함수에게 첫 번째 채널을 닫으라고 알려주기 위한 수단
var CLOSEA = false

var DATA = make(map[int]bool)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func first(min, max int, out chan<- int) { //쓰기전용 채널
	for {
		if CLOSEA {
			close(out)
			return
		}
		out <- random(min, max) //첫 번째 채널에 쓰기
	}
}

func second(out chan<- int, in <-chan int) { //읽기쓰기 모두 가능
	//in채널에서 읽어서 out채널로 보냄
	//그런데 second()함수에서 DATA 맵에 있는 난수를 발견하면 전역 변수인 CLOSEA값을 true로 변경하고 , 더 이상 out 채널로 값을 보내지 않음
	for x := range in {
		fmt.Print(x, " ")
		_, ok := DATA[x]

		if ok {
			CLOSEA = true
		} else {
			DATA[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

func third(in <-chan int) { //읽기전용채널

	var sum int
	sum = 0
	for x2 := range in {
		sum += x2
	}
	fmt.Printf("The sum of the random numbers is %d\n", sum)
}

func main() {
	n1 := 100
	n2 := 200

	rand.Seed(time.Now().Unix())
	A := make(chan int)
	B := make(chan int)

	go first(n1, n2, A)
	go second(B, A)
	third(B)

}
