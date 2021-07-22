package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Epoch time : ", time.Now().Unix())
	//유닉스 에포크 시간 : 1970년 1월 1일 UTC시각으로 경과된 시간을 초 단위로 반환
	t := time.Now()
	fmt.Println(t, t.Format(time.RFC3339)) //Format메소드를 사용하면 time변수를 다른 포맷으로 변환할 수 있다
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())

	time.Sleep(5 * time.Second)
	//time.Second 상수를 이용하면 go 언어에서 1초가 경과된 시간을 표현할 수 있다.
	//10초를 표현하고 싶다면 10 * time.Second를 하면된다.
	t1 := time.Now()
	fmt.Println("Time Difference : ", t1.Sub(t))

	formatT := t.Format("01 January 2006")
	fmt.Println("Format : ", formatT)
	loc, _ := time.LoadLocation("Europe/Paris")
	londonTime := t.In(loc)
	fmt.Println("Paris : ", londonTime)
}
