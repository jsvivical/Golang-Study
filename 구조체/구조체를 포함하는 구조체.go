//내장 타입처럼 포함하는 방식
package main

import "fmt"

type User struct {
	name string
	ID   string
	age  int
}

type VIP struct {
	userInfo User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 23}
	vip := VIP{
		User{"화랑", //name
			"hwarnag", //ID
			40,        //age
		},
		3,   // VIP level
		250, //price
	}

	fmt.Println(user)
	fmt.Println(vip)

}
