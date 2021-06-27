package main

import "fmt"

func PrintValue(v interface{}) {
	switch t := v.(type) /*[인터페이스 변수].(type)을 통해 구체화된 타입으로 변환 */ {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("Not supported type: %T : %v \n", t, t)
	}
}

type Student struct {
	age int
}

func main() {
	PrintValue(10)
	PrintValue(3.14)
	PrintValue("Hello")
	PrintValue(Student{15})

}
