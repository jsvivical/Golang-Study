package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("please give one or more floats")
		os.Exit(1)
	}
	arguments := os.Args // os.Args는 string값을 원소로 갖는 슬라이스로
	//이 슬라이스의 첫번재 원소는 실행 가능한 프로그램의 이름이다.
	min, _ := strconv.ParseFloat(arguments[1], 64) //스트링 입력을 산술 데이터 타입으로 변환
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[1], 64)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Println("Min : ", min)
	fmt.Println("Max : ", max)

}
