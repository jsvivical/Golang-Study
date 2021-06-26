package main

import "fmt"

func main(){
	c := Add(3, 6)
	fmt.Print(c)
}

func Add(a int, b int) int {
	return a + b

}