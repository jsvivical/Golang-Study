package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

/******************************************************************************************************************************************
 * 구조를 모르는 JSON 파일을 읽고 저장한는 방법이다.
 * 이런 JSON 데이터는 GO 구조체가 아닌 GO맵을 이용한다.
******************************************************************************************************************************************/
func main() {
	filename := "noStr.json"
	fileData, err := ioutil.ReadFile(filename)
	//ioutil.ReadFile()은 파일 전체를 한번에 읽음
	if err != nil {
		log.Fatal(err)
	}

	//변수 ParsedData는 맵을 정의하고, 앞에서 읽은 JSON파일의 내용을 여기에 담음
	//string타입의 맵의 키마다 json 속성을 대응시킨다.
	//맵 키에 대한 값은 interface{}이므로, 모든 타입의 값을 가질 수 있다.
	var ParsedData map[string]interface{}
	//json.unmarshal()은 파일의 내용을 ParsedData 맵에 담는데 사용했다.
	json.Unmarshal([]byte(fileData), &ParsedData)

	for key, value := range ParsedData {

		fmt.Printf("[%10v]%20v\n", key, value)
	}

}
