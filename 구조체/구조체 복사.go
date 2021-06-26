package main

import "fmt"

type Student struct {
	age   int
	no    int
	score int
}

func PrintStudent(s Student) {
	fmt.Printf("나이 : %d\n", s.age)
	fmt.Printf("번호 : %d\n", s.no)
	fmt.Printf("점수 : %d\n", s.score)
}

func main() {
	student1 := Student{15, 23, 88}
	student2 := student1

	PrintStudent(student1)
	fmt.Println()
	PrintStudent(student2)
}
