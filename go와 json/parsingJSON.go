package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*****************************************************************************************************************
 * son은 자바스크립트 시스템끼리 간편하고 가볍게 정보를 주고 받게 설계된 굉장히 인기있는 택스트 기반 포맷이다.
 * 그런데 JSON은 애플리케이션의 설정파일에서 데이터를 구조적 형식에 맞게 저장하는 용도로 사용된다.
 * encoding/json 패키지는 go 오브젝트와 JSON문서를 서로 변환하기 위한 Encode()와 Decode()함수를 제공한다.
 *  Encode(), Decode()는 다중 오브젝트나 바이트 스트림에 적용되고 Marshal()/Unmarshal()은 단일 오브젝트에 대해 적용하고
*****************************************************************************************************************/

//JSON데이터를 담을 구조체 변수 정의
//변수이름은 json파일에 있는 것과 같아야함(소대문자까지)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func LoadFromJSON(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}
	in.Close()
	return nil
}

func main() {

	filename := "C:\\Users\\jsviv\\OneDrive\\바탕 화면\\go\\go와 json\\readMe.json"

	var myRecord Record
	err := LoadFromJSON(filename, &myRecord)
	if err == nil {
		fmt.Println(myRecord)
	} else {
		fmt.Println(err)
	}
}
