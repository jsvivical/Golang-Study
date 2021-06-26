package main

import "fmt"

func main() {
	day := "thursday"

	switch day {
	case "monday", "tuesday":
		fmt.Print("월, 화요일은 수업 가는 날입니다.")
	case "wednesdat", "thursday", "friday":
		fmt.Print("수, 목, 금요일은 실습가는 날입니다.")

	}

}
