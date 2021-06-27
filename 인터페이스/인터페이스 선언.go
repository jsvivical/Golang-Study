package main

import "fmt"

type Stringer interface {
	String() string
}

// 덕타이핑 : 타입 선언 시 인터페이스 구현 여부를 명시적으로 나타낼 필요없이 인터페이스에 정의한 메소드 포함 여부만으로 결정하는 방식

type Student struct /*implements Stringer*/ {
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
