// +build ignore
package main

import "fmt"

func Multiple(a, b int) int {
	return a * b
}
func main() {
	var a, b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		result := Multiple(a, b)
		fmt.Println(result)
		fmt.Print()
	}

}
