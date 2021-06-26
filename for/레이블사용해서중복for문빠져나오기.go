package main

import "fmt"

func main() {
	var a, b int = 1, 1

OuterFor: //레이블 정의
	for a = 1; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OuterFor
			}
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)

}
