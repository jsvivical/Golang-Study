package main

import "fmt"

func main() {
	willClose := make(chan int, 10)
	// 채널에 데이터를 쓴다.
	willClose <- -1
	willClose <- 0
	willClose <- 2
	//채널을 비운다.
	<-willClose
	<-willClose
	<-willClose
	close(willClose)
	read := <-willClose
	fmt.Println(read)
}
