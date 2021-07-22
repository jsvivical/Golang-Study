package main

import "fmt"

type Memoized func(int) int

var fibMem Memoized

func Memoize(mf Memoized) Memoized {
	cache := make(map[int]int)
	return func(key int) int {
		if val, found := cache[key]; found {
			return val //이미 결과값이 저장되있는 함수이면 저장된 결과 리턴
		}
		temp := mf(key) //func(int) int 인 피보나치 함수 호출해서 결과값 저장
		cache[key] = temp
		return temp

	}
}

func fib(x int) int {
	if x == 0 {
		return 0
	} else if x <= 2 {
		return 1
	} else {
		return fib(x-2) + fib(x-1)
	}
}

func FibMemoized(n int) int {
	return fibMem(n)
}

func main() {
	fmt.Println()
}
