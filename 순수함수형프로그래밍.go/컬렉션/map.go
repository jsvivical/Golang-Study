package main

import "fmt"

/*
map개요
map은 키에 대응하는 값을 신속히 찾는 해시 테이블을 구현한 자료구조
" map[key타입] value타입 "과 같이 선언
example -> var idMap map[int] string
이 때, 선언된 idMap은 nil값을 가지며, 이를 nil Map이라 부른다.
nil Map에는 어떤 데이터도 쓸 수 없는데, map을 초기화 하기 ㅎ위해
make() 함수를 사용
idMap = make(map[int] string)
make() 함수의 첫 번째 파라미터로 map 키워드와 [키타임] 값타입을 지정하는데, 이때의
이때의 make()함수는 해시테이블 자료구조를 메모리에 생성하고 그 메모리를 가리키는 map value를 리턴한다.
map은 make() 함수를 써서 초기화 할 수도 있지만, 리터럴을 사용해 초기화 할 수도 잇다.
*/

func main() {

	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
	}

	fmt.Println(tickers)

	//Map사용

	var m map[int]string
	m = make(map[int]string)

	m[901] = "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"

	str := m[134]
	fmt.Println(str)

	noDate := m[999] //값이 없으면 nil 혹은 zero리턴
	fmt.Println(noDate)

	delete(m, 777)

	fmt.Println(m)

	//Map 키체크
	//Go에선 map변수[키] 읽기를 수행할 때 2개의 리턴값을 리턴
	//1. 키에 상응하는 값, 2. 그 키가 존재하는지 아닌지를 나타내는 bool값

	v, isExists := tickers["MSFT"]
	if !isExists {
		fmt.Println("No MSFT tickers")
	} else {
		fmt.Println(v)
	}

	//for 루프를 사용한 Map 열거

	myMap := map[string]string{
		"A": "Apple",
		"B": "Banana",
		"C": "Charlie",
		"D": "Delta",
	}

	for key, value := range myMap {
		fmt.Println(key, value)
	}

}
