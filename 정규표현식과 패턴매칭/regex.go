/*************************************************************************************************************************************
 * 패턴 매칭  : 스트링에서 특정한 문자 집합을 찾는 것으로 ,Go 언어에서 핵심적인 역할을 담당한다.
 * 찾으려는 문자집합은 정규표현식과 문법에 맞춰 작성한 검색 패턴으로 표현한다.
 * Go 패키지 중에서 regexp는 정규표현식을 정의하고 패턴 매칭을 수행하는 기능 제공
*************************************************************************************************************************************/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Printf("usage : selectColumn column [<file1> [<file1> [...<fileN>]]\n")
		os.Exit(1)
	}
	temp, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("column value is not n integer : ", temp)
		return
	}
	column := temp
	if column < 0 {
		fmt.Println("Invalid Column number!")
		os.Exit(1)
	}

	for _, filename := range arguments[2:] {
		fmt.Println("\t\t", filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error opening file %s\n", err)
			continue
		}
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')

			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading File %s", err)
			}
			data := strings.Fields(line)
			if len(data) >= column {
				fmt.Println(data[column-1])
			}
		}
	}
}
