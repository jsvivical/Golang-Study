//캡쳐 : 함수 리터럴 외부 변수를 내부 상태로 가져오는 것
//캡쳐는 값 복사가 아닌 참조 형태로 가져오게 되니 주의해야 함

package main

import "fmt"

func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		f[i] = func() { fmt.Println(i) }
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop2")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() { fmt.Println(v) }
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func main() {
	CaptureLoop()
	CaptureLoop2()
}
