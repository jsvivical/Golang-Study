package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.5, 12.5, 53.4, 23.7, 12.1}

	for i, v := range t {
		fmt.Println(i, v) //i 는 인덱스, v는 값
	}

	for _, v := range t {
		fmt.Println(v)
	}

}
