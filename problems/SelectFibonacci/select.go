package main

import "fmt"

func subRoutine(quit, c chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
	//   close(c)
}

func main() {
	c := make(chan int)
	q := make(chan int)
	go subRoutine(q, c)
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	q <- 0
}
