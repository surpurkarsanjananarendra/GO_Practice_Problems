package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Balance struct {
	mu      sync.Mutex
	balance int
}

func (b *Balance) Deposit(amt int, wg *sync.WaitGroup) {
	defer wg.Done()

	b.mu.Lock()
	defer b.mu.Unlock()
	for i := 0; i < 5; i++ {
		fmt.Printf("Person %d :\n", i+1)
		fmt.Printf("Deposit Amount: %d\n\n", amt)
		b.balance += amt
	}

}

func main() {
	person := &Balance{balance: 0}

	fmt.Printf("Current Balance: %d\n", person.balance)

	wg.Add(1)
	go person.Deposit(1, &wg)

	wg.Wait()
	fmt.Println("All Amount Deposited Successfully...")
	fmt.Printf("Your Balance after Deposit: %d \n\n", person.balance)
}
