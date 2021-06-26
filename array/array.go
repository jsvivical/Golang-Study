package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 34.3, 12.5, 25.5, 75.3}

	for i := 0; i < 5; i++ {
		fmt.Printf("%-5.2f , ", t[i])
	}
}
