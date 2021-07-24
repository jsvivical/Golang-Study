package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

/******************************************************************************************************************************************
 * Go는 xml을 지원한다.
 * XML은 HTML과 비슷하지만 그보다 더 고급인 마크업언어이다.
 * 이 프로그램은 디스크에 저장된 JSON파일을 읽어 내용을 변경한 후 XML로 변환해서 화면에 출력하는 프로그램이다.
*****************************************************************************************************************************************/

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func loadFromJSON(filename string, key interface{}) error {
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
	var filename string = "redMe.json"

	var myRecord Record
	err := loadFromJSON(filename, &myRecord)
	if err == nil {
		fmt.Println("JSON : ", myRecord)
	} else {
		fmt.Println(err)
	}
	myRecord.Name = "Dimitris"

	xmlData, _ := xml.MarshalIndent(myRecord, "", "	")
	xmlData = []byte(xml.Header + string(xmlData))
	fmt.Println("\nxmlData : ", string(xmlData))

	data := &Record{}
	err = xml.Unmarshal(xmlData, data)
	if err != nil {
		fmt.Println("error Unmarshaling to json")
		return
	}

	result, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error mashaling to JSON")
		return
	}
	_ = json.Unmarshal([]byte(result), &myRecord)
	fmt.Println("\nJSON : ", myRecord)

}
