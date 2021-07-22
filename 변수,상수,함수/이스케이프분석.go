package main

type user struct {
	name  string
	email string
}

func stayOnStack() user {
	u := user{
		name:  "Hoanh An",
		email: "hoangan101@gmail.com",
	}
	//변수 u는 stayOnStack 함수에서 벗어나지 못한다.
	//컴파일할 때 u의 크기를 알 수 있기에 컴파일러는 u를 스택 프레임에 저장한다

	return u
}

func escapeToHeap() *user {
	u := user{
		name:  "Hoanh An",
		email: "hoanhan101@gmail.com",
	}
	//escapeToHeap에서는 변수가 함수 바깥으로 나온다. 
	
	return &u
}

func main() {

}
