package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Mutex: Mutual exclusion

https://en.wikipedia.org/wiki/Mutual_exclusion
*/

type Account struct {
	balance int
	lock    sync.Mutex
}

// return the balance of an account
func (a *Account) GetBalance() int {
	a.lock.Lock()
	// defer the unlock to the end of the function
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) Withdraw(v int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance {
		fmt.Println("Not enough money in account")
	} else {
		fmt.Printf("%d withdrawn : Balance : %d\n", v, a.balance)
		a.balance -= v
	}

}

func main() {
	var acc1 Account
	acc1.balance = 100

	fmt.Println("Balance:", acc1.GetBalance())

	for i := 0; i < 12; i++ {
		go acc1.Withdraw(10) // goroutine
	}
	time.Sleep(2 * time.Second)
}
