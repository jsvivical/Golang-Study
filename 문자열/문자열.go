package main

import "fmt"

func main() {
	// "" (큰따옴표)와 ``(백 쿼트, 그레이브)는  그 쓰임이 다름
	// ``로 문자열을 묶으면 문자열 안의 특수 문자가  일반 문자처럼 처리됨

	a := "Hello\nWorld"
	b := `Hello\nWorld`

	fmt.Println(a)
	fmt.Println(b)

	c := "죽는 날까지 하늘을 우러러\n한 점 부끄럼 없기를 \n잎새에 이는 바람에도\n나는 괴로워했다.\n"

	d :=
		`죽는 날까지 하늘을 우러러
한 점 부끄럼 없기를
잎새에 이는 바람에도 
나는 괴로워했다.`

	fmt.Println(c)
	fmt.Println()
	fmt.Println()
	fmt.Println(d)
}
