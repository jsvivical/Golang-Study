package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Element struct {
	name    string
	surname string
	id      string
}

//DATA는 맵을 표현
//키 값은 string, 값은 Element 타입으로 지정
var DATA = make(map[string]Element)

//DATA에 추가
func ADD(k string, n Element) bool {
	if k == "" {
		return false
	} //공백문자는 키 값이 될 수 없음

	if LOOKUP(k) == nil {
		DATA[k] = n
		return true
	} //기존에 k값을 가지는 키가 없음 -> 새로운 값

	return false //이미 존재하는 키 값을 받은 경우
}

//주의할 점
//새로운 원소를 스토어에 추가할  때 Element 구조체를 모두 채울 만큼 충분히 값을 지정하지 않으면 ADD함수가 제대로 실행되지 않는다.
//이 예제에서는 구조체에서 빠진 부분을 공백 스트링으로 채운다.
//하지만 기존 키를 추가하면 기존 키- 값을 변경하는 것이 아니라 에러메시지를 출력한다.

func DELETE(k string) bool {
	if LOOKUP(k) != nil { //이미 존재하는 값
		delete(DATA, k)
		return true
	}
	return false //존재하지 않는 값을 삭제할 순 없음
}

func LOOKUP(k string) *Element {
	_, ok := DATA[k]
	if ok { //기존에 키가 존재하면 그 값의 주소 반환
		n := DATA[k]
		return &n
	} else { //주어진 값의 키가 없으면 nil 반환
		return nil
	}
}

func CHANGE(k string, n Element) bool {
	DATA[k] = n
	return true
}
func PRINT() {
	for k, d := range DATA {
		fmt.Printf("[%v] %v\n", k, d)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		tokens := strings.Fields(text)

		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 2:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 3:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 4:
			tokens = append(tokens, "")
		}
		//첫번째 switch문 이유
		//tokens슬라이스 원소의 개수가 최소한 다섯개가 되도록 설정

		switch tokens[0] { //tokens[0]은 PRINT, ADD 등의 명령어임
		case "PRINT":
			PRINT()
		case "STOP":
			return
		case "DELETE":
			if !DELETE(tokens[1]) {
				fmt.Println("Delete operation failed!\n")
			}
		case "ADD":
			n := Element{tokens[2], tokens[3], tokens[4]}
			if !ADD(tokens[1], n) {
				fmt.Println("Add operation failed\n")
			}

		case "LOOKUP":
			n := LOOKUP(tokens[1])
			if n != nil {
				fmt.Printf("%v\n", *n)
			}

		case "CHANGE":
			n := Element{
				name:    tokens[2],
				surname: tokens[3],
				id:      tokens[4],
			}
			if !CHANGE(tokens[1], n) {
				fmt.Println("Update operation failed")
			}
		default:
			fmt.Println("unknown command : please try again!\n")
		}
	}
}
