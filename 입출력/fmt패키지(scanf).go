package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Println("scanf")
	n, err := fmt.Scanf("%d %d", &a, &b)

	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
	fmt.Println("scanln")
	fmt.Scanln()
	n2, err2 := fmt.Scanln(&a, &b)

	if err2 != nil {
		fmt.Println(n2, err)
	} else {
		fmt.Println(n2, a, b)
	}

}
