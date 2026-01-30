package main

import (
	"fmt"
	"sync"
)

func even(e []int) {
	var sum int
	for _, val := range e {
		if val%2 == 0 {
			sum += val
		}
	}
	fmt.Printf("Sum of Evens: %d\n", sum)
}

func odd(o []int) {
	var sum int
	for _, val := range o {
		if val%2 != 0 {
			sum += val
		}
	}
	fmt.Printf("Sum of Odds: %d\n", sum)
}

func main() {
	var wg sync.WaitGroup

	e := [5]int{1, 2, 3, 4, 5}

	wg.Go(func() {
		even(e[:])
		odd(e[:])
	})

	wg.Wait()
	fmt.Println("Successfull Execution...")
}
