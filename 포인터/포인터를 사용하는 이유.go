package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg *Data) {
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data

	ChangeData(&data)

	fmt.Println(data)
}
