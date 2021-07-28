package main

import "fmt"

/************************************************************************************************************************************
 * 타입 메서드 : 특수한 수신자 인수를 받는 함수
 * 수신자의 작동 과정을 OOP용어로 표현하면 객체에 메시지를 보낸다고 한다.
 *
************************************************************************************************************************************/

type twoInts struct {
	x int64
	y int64
}

func regularFunction(a, b twoInts) twoInts {
	temp := twoInts{x: a.x + b.x, y: a.y + b.y}
	return temp
}

func (a twoInts) method(b twoInts) twoInts {
	temp := twoInts{x: a.x + b.x, y: a.y + b.y}
	return temp
}

func main() {
	i := twoInts{x: 1, y: 2}
	j := twoInts{x: -5, y: -2}
	fmt.Println(regularFunction(i, j))
	fmt.Println(i.method(j))
}
