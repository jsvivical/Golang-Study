package main

import "fmt"

func main() {
	temp := 18

	switch true {
	case temp < 10 , temp > 30:
		fmt.Print("It is not a good day to go outside")

	case temp >= 10 && temp < 20:
		fmt.Print("It's bit cold, so put your outer on when you go outside")

	case temp >= 20 && temp < 25:
		fmt.Print("Perfect weather to go outside")
	default:
		fmt.Print("It is warm")

	}

}
