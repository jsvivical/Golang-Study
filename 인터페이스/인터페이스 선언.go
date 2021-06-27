package main

import "fmt"

type Stringer interface {
	String() string
}
type Student struct {
	name string
	age  int
}

func (s Student) String() string /*메소드*/ {
	return fmt.Sprintf("안녕 나는 %d살 %s라고 해\n", s.age, s.name)
}

func main() {
	student := Student{"철수", 12}

	var stringer Stringer

	stringer = student
	fmt.Printf("%s\n", stringer.String())
}
