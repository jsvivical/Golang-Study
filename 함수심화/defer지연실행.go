//함수가 종료되기 직전에 실행해야 하는 코드가 있는 경우
//사용법 : defer 명령문

package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("Test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer fmt.Println("반드시 호출됩니다.")
	defer f.Close()
	defer fmt.Println("파일을 닫았습니다.")

	fmt.Println("파일에 Hello wolrd를 씁니다.")
	fmt.Fprintln(f, "Hello Wrold")
}
