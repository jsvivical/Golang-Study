package main

import "fmt"

type House struct {
	address string
	size    int
	price   float32
	Type    string
}

func main() {
	var house House
	house.address = "서울시 강동구 ..."
	house.size = 28
	house.price = 9.8
	house.Type = "아파트"

	fmt.Printf("주소 : %-s\n", house.address)
	fmt.Printf("크기 : %-d\n", house.size)
	fmt.Printf("가격 : %-.2f\n", house.price)
	fmt.Printf("타입 : %-s\n ", house.Type)

}
