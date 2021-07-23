package main

import (
	"encoding/json"
	"fmt"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func main() {
	myRecord := Record{
		Name:    "Mihails",
		Surname: "Tsoukalos",
		Tel: []Telephone{
			Telephone{Mobile: true, Number: "1234- 567"},
			Telephone{Mobile: true, Number: "1234-abcd"},
			Telephone{Mobile: false, Number: "abcc - 567"},
		},
	}

	rec, err := json.Marshal(&myRecord)
	//Marshal()은 포인터 변수를 입력받아 json포맷으로 변환한다.
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(rec))

	var unRec Record
	err1 := json.Unmarshal(rec, &unRec)
	//Unmarshal()은 JSON을 입력받아 go구조체로 변환한다.
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(unRec)
}
