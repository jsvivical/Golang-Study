package main

import "fmt"

func main() {
	a := 3
	//break키워드는 사용하나 마나  하나의 케이스를 살피고 빠져나옴

	switch a {
	case 1:
		fmt.Println("a == 1")
		break
	case 2:
		fmt.Println("a == 2")
		break
	case 3:
		fmt.Println("a == 3")
	case 4:
		fmt.Println("a == 4")
	default:
		fmt.Println("a > 4")
	}
	//fallthrough는 다음 case까지 실행됨
	switch a {
	case 1:
		fmt.Println("a == 1")
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
		fallthrough //다음 케이스 문 실행
	case 4:
		fmt.Println("a == 4")
		fallthrough //다음 케이스 문 실행
	default:
		fmt.Println("a > 4")
	}

}
