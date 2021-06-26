package main

import "fmt"

type House struct {
	address string
	size    int
	price   float64
	Type    string
}

func main() {
	//초깃값을 생략하면 모든 필드가 기본값으로 초기화됨

	var house1 House
	fmt.Println(house1)

	//모든 필드 초기화 {} 사용
	var house2 House = House{
		"서울시 강동구",
		28,
		9.80,
		"아파트", //세로로 펼쳐서 초기화 할 때는 반드시 마지막 요소 다음에 ' ,  ' 를 넣어줘야 함.
	}
	fmt.Println(house2)

	//d일부 필드 초기화
	house3 := House{size: 28, Type: "아파트"}
	fmt.Println(house3)

}
