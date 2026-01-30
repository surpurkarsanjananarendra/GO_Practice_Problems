package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func even(e []int, wg *sync.WaitGroup) {
	defer wg.Done()
	var sum int
	for _, val := range e {
		if val%2 == 0 {
			sum += val
		}
	}
	fmt.Printf("Sum of Evens: %d\n", sum)
}

func odd(o []int, wg *sync.WaitGroup) {
	defer wg.Done()
	var sum int
	for _, val := range o {
		if val%2 != 0 {
			sum += val
		}
	}
	fmt.Printf("Sum of Odds: %d\n", sum)
}

func main() {
	e := [5]int{1, 2, 3, 4, 5}

	wg.Add(2)

	go even(e[:], &wg)
	go odd(e[:], &wg)

	wg.Wait()
	fmt.Println("Successfull Execution...")
}
