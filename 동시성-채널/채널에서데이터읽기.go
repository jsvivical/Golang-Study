package main

import (
	"fmt"
	"time"
)

/*****************************************************************************************************************************************
*****************************************************************************************************************************************/

func writeToChannel(c chan int, x int) {
	fmt.Println("1", x)
	c <- x
	close(c)
	fmt.Println("2", x)
}

func main() {
	c := make(chan int)
	go writeToChannel(c, 10)
	time.Sleep(1 * time.Second)
	fmt.Println("Read : ", <-c) //<-c라는 표현식을 통해 채널c에서 데이터를 읽음
	//채널에서 읽은 데이터를 변수에 저장하려면 k := <- c라고 하면 된다.
	time.Sleep(1 * time.Second) //채널에서 데이터를 읽는 시간 지정

	//아래는 현재 채널이 열렸는지 확인하는 방법
	v, ok := <-c
	if ok == true {
		fmt.Println("Channel is open!\n")
		fmt.Println(v)
	} else {
		fmt.Println("Channel is closed\n")
	}
}
