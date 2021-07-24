/**
 * 이 프로그램은 여러 가지 옵션을 받는데, 값이 여러 개일 때, 각각을 콤마(,)로 구분
 * 이 프로그램은 flag.Var()함수를 이용해 flag.Value 인터페이스를 만족하는 모든 종류의 생성
 * */

package main

import (
	"flag"
	"fmt"
	"strings"
)

type Value interface {
	String() string
	Set() error
}

type NamesFlag struct {
	Names []string
}

func (s *NamesFlag) GetNames() []string {
	return s.Names
}

func (s *NamesFlag) String() string {
	return fmt.Sprint(s.Names)
}

func (s *NamesFlag) Set(v string) error {
	if len(s.Names) > 0 {
		return fmt.Errorf("cannot use names flag more than once!")
	}

	names := strings.Split(v, ",")
	for _, item := range names {
		s.Names = append(s.Names, item)
	}
	return nil
}

func main() {
	var manyNames NamesFlag
	minusK := flag.Int("k", 0, "An int")
	minusO := flag.String("o", "Mihails", "The name")
	flag.Var(&manyNames, "names", "Comma-seperated list")

	flag.Parse()
	fmt.Println("-k : ", *minusK)
	fmt.Println("-o : ", *minusO)

	for i, item := range manyNames.GetNames() {
		fmt.Println(i, item)
	}

	fmt.Println("Remaining command-line arguments :")
	for index, val := range flag.Args() {
		fmt.Println(index, " : ", val)
	}

}
