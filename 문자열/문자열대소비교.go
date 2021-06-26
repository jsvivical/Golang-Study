package main

import "fmt"

func main() {
	str1 := "BBB"
	str2 := "aaaaAAAA"
	str3 := "BBAD"
	str4 := "ZZZ"
	str5 := "git 연습용"
	fmt.Printf("%s > %s : %v\n", str1, str2, str1 > str2)
	fmt.Printf("%s < %s : %v\n", str1, str3, str1 < str3)
	fmt.Printf("%s <= %s : %v\n", str1, str4, str1 <= str4)
	fmt.Println(str5)
}
