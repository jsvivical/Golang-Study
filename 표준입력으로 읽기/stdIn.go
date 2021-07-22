package main

import (
	"bufio"
	"fmt"
	"os"
)

//파일 입출력 기능은 주로 bufio패키지에 있는 것을 사용하지만, os 패키지에 유용한 함수들이 많아 거의 모든 패키지에 os패키지를 추가한다.
//그 중에서도 특히 go프로그램의 커맨드라인 인수에 접근하는데 많이 사용한다.

//공식문서에 따르면 os패키지는 OS연산을 수행하는 함수를 제공한다. 파일을 생성하고, 삭제하고, 파일과 디렉토리의 이름을 바꾸는 것뿐만 아니라
//유닉스 접근 권한을 비롯한 파일과 디렉터리에 대하 다양한 속성을 확인하는 함수도 있다.
//os 패키지의 가장 큰 장점은 플랫폼 독립적이라는 것이다.
//즉. 유닉스나 윈도우나 똑같이 작동함

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}
