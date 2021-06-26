package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for {
		time.Sleep(time.Minute / 12) //5초 동안 실행을 멈춤
		i += 5
		fmt.Println(i)
	}
}
