package main

import (
	"fmt"
	"math"
	"myInterface"
)

type square struct {
	x float64
}
type circle struct {
	r float64
}

func (s square) Area() float64 {
	return s.x * s.x
}

func (s square) Perimeter() float64 {
	return 4 * s.x
}

func (c circle) Area() float64 {
	return c.r * c.r * math.Pi
}

func (c circle) Perimeter() float64 {
	return 2 * c.r * math.Pi
}

func Calculate(x myInterface.Shape) {
	_, ok := x.(circle)
	if ok {
		fmt.Println("Is a circle")
	}

	v, ok := x.(square)
	if ok {
		fmt.Println("is a square : ", v)
	}

	fmt.Println(x.Area())
	fmt.Println(x.Perimeter())

}

func main() {
	x := square{x: 10}
	fmt.Println("perimeter : ", x.Perimeter())
	Calculate(x)
	y := circle{r: 5}
	Calculate(y)
}
