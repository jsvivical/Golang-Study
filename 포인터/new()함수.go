package main

import "fmt"

type Data struct {
	num1 int
	num2 int
	num3 int
}

func main() {
	p1 := &Data{}
	var p2 = new(Data)

	p2.num1 = 1
	p2.num2 = 2
	p2.num3 = 3

	fmt.Println(*p1)
	fmt.Println(*p2)

}
