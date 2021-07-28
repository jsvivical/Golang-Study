package main

import "fmt"

type square struct {
	x float64
}

type circle struct {
	r float64
}

type rectangle struct {
	x float64
	y float64
}

func tellInterface(x interface{}) {
	switch v := x.(type) {
	case square:
		fmt.Println("this is a square")
	case circle:
		fmt.Printf("%v  is circle", v)
	case rectangle:
		fmt.Println("this is a rectangle")
	default:
		fmt.Printf("Unknown type %T\n", v)
	}
}

func main() {
	x := circle{r: 10}
	tellInterface(x)
	y := rectangle{x: 4, y: 1}
	tellInterface(y)
	z := square{x: 4}
	tellInterface(z)
	tellInterface(10)
}
