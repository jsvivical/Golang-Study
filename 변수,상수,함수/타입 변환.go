package main

import "fmt"

func main() {
	a := 3
	b := 3.5
	var c int = int(b)
	d := float64(a * c)

	var e int64 = 7
	f := int64(b) * e

	fmt.Print(a, b, c, d, e,f)

}
