package main

import "fmt"

func function1(out <-chan int, i int) { //읽기전용채널
	fmt.Println(i)
	out <- i //읽기전용채널에 쓰려고하면 컴파일러가 오류를 출력

	close(out)
}

func function2(out chan int, i int) {
	fmt.Println(i)
	out <- i

	close(out)

}
