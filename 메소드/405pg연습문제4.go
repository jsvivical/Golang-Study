package main

import (
	"fmt"
	"time"
)

type Courier struct {
	name string
}
type Product struct {
	name  string
	price int
	id    int
}
type Parcel struct {
	pdt           *Product
	ShippedTime   time.Time
	DeliveredTime time.Time
}

func (courier *Courier) SendProduct(product *Product) *Parcel {
	var parcel *Parcel = new(Parcel)
	parcel.pdt = product
	parcel.ShippedTime = time.Now()

	return parcel
}

func (parcel *Parcel) Delivered() *Product {
	parcel.DeliveredTime = time.Now()

	return parcel.pdt
}

func main() {

	var courier Courier
	courier.name = "한진택배"

	var product Product
	product.id = 1011201
	product.name = "도서"
	product.price = 11900

	fmt.Println("배송을 시작합니다.")
	var parcel = courier.SendProduct(&product)
	fmt.Println(parcel.ShippedTime)
	fmt.Println("배송이 도착")
	parcel.Delivered()
	fmt.Println(parcel.DeliveredTime)

}
