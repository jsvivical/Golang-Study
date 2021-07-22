package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

func main() {
	var wg sync.WaitGroup

	account := &Account{0}

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func DepositAndWithdraw(account *Account) {
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value : -1000"))
	}
	account.Balance += 1000 //오류가 나는 원인 : 한 고루틴에서 값을 변경할 때, 다른 고루틴에서도 값의 변경이 일어남
	fmt.Println("입금 : ", account.Balance)

	time.Sleep(1000 * time.Millisecond)
	account.Balance -= 1000
	fmt.Println("출금 : ", account.Balance)

}
