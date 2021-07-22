//뮤텍스는 동시성 프로그래밍의 문제를 해결할 수 있지만, 또 다른 문제가 발생할 수 있다
//1. 동시성 프로그래밍으로 얻을 수 있는 성능 향상을 얻을 수 없다.
//성능을 향상시키려고 동시성 프로그램을 구현했지만 성능향상을 얻지 못하는 아이러니한 문제가 생김

//2. 데드락이 발생할 수 있음
//데드락은 프로그램이 완전히 멈춰 버리는 문제

//아래 프로그래밍은 데드락이 발생함

//테이블 위에 수저와 포크가 하나씩 있고, 식사는 수저,포크 모두가 있어야 가능하다. A,B 각각 수저,포크를 집었음
//그래서 A,B는 식사를 하지 위해 포크,수저를 집으려하지만 테이블 위에 없고, 서로 양보하지 않아서 둘 다 식사를 하지못함.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)

	fork := &sync.Mutex{}
	spoon := &sync.Mutex{} //포크와 수저 뮤텍스

	go diningProblem("A", fork, spoon, "포크", "수저")
	go diningProblem("B", spoon, fork, "수저", "포크")
	wg.Wait()

}

func diningProblem(name string, first, second *sync.Mutex, firstname, secondname string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s가 밥을 먹으려 합니다.\n", name)
		first.Lock()
		fmt.Printf("%s %s 획득\n", name, firstname)
		second.Lock()
		fmt.Printf("%s %s 획득\n", name, secondname)

		fmt.Printf("%s가 식사를 합니다.\n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	wg.Done()
}
