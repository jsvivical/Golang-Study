package main

type myStructure struct {
	Name    string
	Surname string
	height  int32
}

func createStruct(n, s string, h int32) *myStructure {
	if h > 300 {
		h = 0
	}
	return &myStructure{n, s, h}
}

func returnStructure(n, s string, h int32) myStructure {
	if h > 300 {
		h = 0
	}
	return myStructure{n, s, h}
}

//위의 두 함수의 기능은 같다.

func main() {

}
