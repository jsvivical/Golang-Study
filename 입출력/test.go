package main

func main() {
	var integer int = 34
	var str string = "HELLO WORLD!"
	var i int = 0
	var limit = 10
	for i = 0; i < 10; i++ {
		println(i)
	}
	i = 0

	for i < 10 {
		print(10 - i)
		i++
	}
	println()
	i = 0
	for i != limit {
		print(i)
		i++
	}
	println("Hello World")
	println("Bye World")
	println(integer)
	println(str)
}
