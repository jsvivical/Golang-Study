package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	age int16
}

func (s *Student) String() string {
	return fmt.Sprintf("Student age : %d", s.age)
}

func PrintAge(stringer Stringer) {
	s := stringer.(*Student) //s는 지역변수 -> 새로 할당받는거임
	fmt.Printf("age : %d\n", s.age)
}

func main() {
	s := &Student{15} //s는 포인터변수
	fmt.Println(s.String())
	PrintAge(s)

}
