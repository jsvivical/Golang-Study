/**************************************************************************************************************************************
 * viper는 다양한 옵션을 제공하는 강력한 Go 패키지
 * viper패키지는 일정한 패턴을 따름
 *1. viper를 초기화한 후 원하는 원소를 정의
 * 2. 이런 원소를 가져온 후 각각의 값을 읽어 정의한다.
 *
 *viper패키지는 flag패키지의 기능을 완전히 포함한다.
 *
 *JSON,YAML,TOML, HCL, 자바, 프로퍼티 포맷 등과 같이 정해진 설정 파일을 사용할 때 파싱 작업을 viper가 모두 처리
 *따라서 코드를 작성하고 디버깅하는 데 드는 수고를 크게 덜어준다.
 *
 *viper는 Go 구조체에서 값을 가져오거나 저장하는 기능도 제공, 이때 Go구조체의 필드와 설정 파일의 키를 매치해야 한다.
 *viper 홈페이지 (https://github.com/spf13/viper)
**************************************************************************************************************************************/
package main

import (
	"fmt"
	
	"github.com/spf13/viper"
)

func main() {
	viper.BindEnv("GOMAXPROCS")
	val := viper.Get("GOMAXPROCS")
	fmt.Println("GOMAXPROCS : ", val)

	viper.Set("GOMAXPROCS", 10)
	val = viper.Get("GOMAXPROCS")
	fmt.Println("GOMAXPROCS:", val)

	viper.BindEnv("NEW_VARIABLE")
	val = viper.Get("NEW_VARIABLE")
	if val == nil {
		fmt.Println("NEW_VARIABLE not defined")
		return
	}
	fmt.Println(val)

}
