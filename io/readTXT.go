/**************************************************************************************************************************************
 *io.Reader, io.Writer 인터페이스를 따르려면 Read(),Write()를 구현해야한다.

 *버퍼를 이용한 파일 입출력과 버퍼를 이용하지 않는 파일 입출력
 *	버퍼를 이용한 파일 입력과 출력은 데이터를 읽거나 쓰기 전에 잠시 버퍼에 저장한다
 *	따라서 파일을 한 바이트 단위로 읽지 않고 한 번에 여러 바이트를 읽을 수 있다.
 * 	버퍼를 사용해야 하는 경우, 사용해야 하지 않는 경우
 * 		1.사용해야 하는 경우 : 중요한 데이터를 다루는 경우 ->버퍼를 거치는 동안, 안 중요해지거나 전원이 꺼진 사이에 데이터를 잃을 수도 있음
 *
 * bufio 패키지
 *
 *
**************************************************************************************************************************************/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func lineByLine(filename string) error {
	var err error

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Error reading file %s", err)
			break
		}
		fmt.Print(line)
	}
	return nil
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: byLine<file1>[<file2>...]\n")
		return
	}

	for _, file := range flag.Args() {
		err := lineByLine(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
