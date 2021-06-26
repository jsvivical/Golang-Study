package main

import "fmt"

const (
	Red   int = iota
	Blue  int = iota
	Green int = iota
)

const (
	C1 uint = iota + 1
	C2
	C3
)

//타입 없는 상수

const PI = 3.14
const FloatPI float64 = 3.14

func main() {
	fmt.Print(Red, Blue, Green, C1, C2, C3)

}
