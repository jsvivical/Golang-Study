package main

import "fmt"

func main() {
	//첫번째방법
	//1. 슬라이스 맨 뒤에 요소 추가
	//2.맨 뒤 값부터 삽입하려는 위치까지 한 칸씩 뒤로 밀어줌
	//3.삽입하는 위치의 값을 바꿔줍니다.

	slice := []int{1, 2, 3, 4, 5, 6} //len(slice) == 6

	idx := 2                                  //추가하려는 인덱스
	slice1 := append(slice, 0)                //len(slice) == 7
	for i := len(slice1) - 2; i >= idx; i-- { //i는 제일 끝의 인덱스(5)부터 시작
		slice1[i+1] = slice1[i]
	}
	slice1[idx] = 100
	fmt.Println(slice1)
	//두번째방법 append()함수이용
	slice2 := slice
	slice2 = append(slice2[:idx], append([]int{100}, slice2[idx:]...)...)
	fmt.Println(slice2)
	//세번째 방법 : 두번째에서 불필요한 메모리 사용이 없도록 코드 개선
	slice3 := append(slice, 0)
	copy(slice3[idx+1:], slice3[idx:])
	slice3[idx] = 100
	fmt.Println(slice3)
}
