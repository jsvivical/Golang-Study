//내장 타입처럼 포함하는 방식
package main

import "fmt"

type User struct {
	name  string
	ID    string
	age   int
	level int
}

type VIP struct {
	userInfo User
	level    int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 23, 10}
	vip := VIP{
		User{"화랑", //name
			"hwarnag", //ID
			40,
			10, //age
		},
		3,   // VIP level
		250, //price
	}

	//"화랑"의 User level에 접근하려면
	vip.userInfo.level = 30

	fmt.Println(user)
	fmt.Println(vip)

}
